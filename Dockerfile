FROM golang:1.21-alpine AS builder

WORKDIR /workspace

RUN apk add --update --no-cache git && rm -rf /var/cache/apk/*
COPY go.mod go.sum /workspace/
RUN go mod download
COPY cmd /workspace/cmd
COPY internal /workspace/internal
RUN go build -o citizen ./cmd/citizen

FROM alpine
RUN apk add --update --no-cache ca-certificates tzdata && rm -rf /var/cache/apk/*
COPY --from=builder /workspace/citizen /usr/local/bin/citizen
CMD [ "/usr/local/bin/citizen", "--server-stream-delay=500ms" ]
