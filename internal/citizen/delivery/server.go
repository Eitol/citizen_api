package delivery

import (
	"connectrpc.com/grpchealth"
	"context"
	"errors"
	"github.com/Eitol/citizen_api/internal/citizen/domain"
	"github.com/Eitol/citizen_api/internal/citizen/repositories/citizenrepo"
	"github.com/Eitol/citizen_api/internal/citizen/usecases"
	"github.com/Eitol/citizen_api/pkg/citizendb/cl"
	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Eitol/citizen_api/internal/citizen/delivery/handlers"
	"github.com/Eitol/citizen_api/internal/gen/go/citizen/api/v1/apiv1connect"
)

type Server struct {
	mux *http.ServeMux
}

func (s *Server) createHandler(serviceName string, f func() (string, http.Handler)) {
	s.mux.Handle(f())
	s.mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(serviceName),
	))
	s.mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(serviceName),
	))
	s.mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(serviceName),
	))
}

func createCitizenRepository() domain.CitizenRepository {
	startTime := time.Now()
	log.Println("CL DB: loading")
	clDB, err := cl.NewDB("scripts/assets/citizen/cl_unified_person_list.gob")
	if err != nil {
		log.Fatalf("error creating cl db: %v", err)
	}
	log.Printf("CL DB: loaded in %v\n", time.Since(startTime))

	startTime = time.Now()
	log.Println("Name list: loading")
	idVsNameDB, err := names.LoadIDVsNameDB("scripts/assets/citizen/names_list.gob")
	if err != nil {
		log.Fatalf("error loading id vs name db: %v", err)
	}
	log.Printf("Name list: loaded in %v\n", time.Since(startTime))

	startTime = time.Now()
	log.Println("VE DB: loading")
	veDB, err := ve.NewCitizenDB(
		"scripts/assets/citizen/ve_optimized_names_list.gob",
		"scripts/assets/citizen/ve_optimized_locations.gob",
		"scripts/assets/citizen/ve_optimized_names_unique.gob",
		idVsNameDB,
	)
	log.Printf("VE DB: loaded in %v\n", time.Since(startTime))
	if err != nil {
		log.Fatalf("error creating ve db: %v", err)
	}
	runtime.GC()
	runtime.GC()
	runtime.GC()
	return citizenrepo.NewMultiCountryCitizenRepository(
		clDB,
		veDB,
	)
}

func (s *Server) Run() {
	s.mux = http.NewServeMux()
	repo := createCitizenRepository()
	uc := usecases.NewFindByIDUC(repo)
	handler := handlers.NewCitizenHandler(uc)
	s.createHandler(apiv1connect.CitizenServiceName, func() (string, http.Handler) {
		return apiv1connect.NewCitizenServiceHandler(
			handler,
		)
	})

	addr := "localhost:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			newCORS().Handler(s.mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err) //nolint:gocritic
	}
}
