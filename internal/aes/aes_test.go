package aes

import (
	"homepass/internal/key"
	"testing"
)

func TestAes(t *testing.T) {
	data := "This is encrypted string"
	salt := []byte("asdfasdf")
	pass := "1q2w3e4r."
	key, err := key.GetKey(salt, pass)
	if err != nil {
		t.Errorf("GetKey error: %v", err)
	}
	want := "3d874b2ff6e8cbfbc73e8dd95f4c1ad5372ef5ceeee1cf43"
	got, err := Encrypt(data, key)

	if err != nil {
		t.Errorf("Encrypt error: %v", err)
	}

	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}

	got, err = Decrypt(want, key)

	if data != got {
		t.Errorf("got %s, want %s", got, data)
	}
}
