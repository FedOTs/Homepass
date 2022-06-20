package key

import (
	"bytes"
	"testing"
)

func TestKey(t *testing.T) {
	salt := []byte("asdfasdf")
	pass := "1q2w3e4r."
	got, err := GetKey(salt, pass)
	if err != nil {
		t.Errorf("GetKey error %v", err)
	}
	want := []byte{44, 218, 6, 170, 177, 116, 160, 204, 165, 95, 253, 73, 16, 229, 162, 205, 137, 104, 61, 180, 210, 188, 64, 92, 23, 234, 160, 248, 207, 212, 150, 185}
	res := bytes.Compare(got, want)
	if res != 0 {
		t.Errorf("GetKey got %v, want %v", got, want)
	}

}
