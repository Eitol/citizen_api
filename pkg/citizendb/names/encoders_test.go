package names

import (
	"reflect"
	"testing"
)

var m = map[string]uint32{
	"Juan":      1,
	"Perez":     2,
	"Rodriguez": 3,
	"Maria":     4,
	"Martinez":  5,
	"Jose":      6,
	"David":     7,
}

var l = []string{
	"Juan",      // 1,
	"Perez",     // 2,
	"Rodriguez", // 3,
	"Maria",     // 4,
	"Martinez",  // 5,
	"Jose",      // 6,
	"David",     // 7,
}

func TestEncodeNamesIn11Bytes(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		want      [4]uint32
	}{
		{
			name:      "TestEncodeNamesIn11Bytes1",
			inputName: "Juan David Rodriguez Martinez",
			want:      [4]uint32{1, 7, 3, 5},
		},
		{
			name:      "TestEncodeNamesIn11Bytes2",
			inputName: "Jose Juan Perez",
			want:      [4]uint32{6, 1, 2, 0},
		},
		{
			name:      "TestEncodeNamesIn11Bytes3",
			inputName: "Maria Rodriguez",
			want:      [4]uint32{4, 3, 0, 0},
		},
		{
			name:      "TestEncodeNamesIn11Bytes4",
			inputName: "Jose",
			want:      [4]uint32{6, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeNamesIn11Bytes(tt.inputName, m)
			want11Bytes := encode4IntsIn11Bytes(tt.want)
			if !reflect.DeepEqual(got, want11Bytes) {
				t.Errorf("EncodeNamesIn11Bytes() = %v, want %v", got, tt.want)
			}

			gotDecoded := DecodeNamesFrom11Bytes(got, l)
			if gotDecoded != tt.inputName {
				t.Errorf("DecodeNamesFrom11Bytes() = %v, want %v", gotDecoded, tt.inputName)
			}
		})
	}
}
