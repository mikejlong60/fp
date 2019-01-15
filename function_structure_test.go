package fp

import (
	"fmt"
	"testing"
)

func TestEncryptedFunctions(t *testing.T) {
	cases := []struct {
		expected string
	}{
		{"fred"},
	}
	for _, c := range cases {

		var funcs = EncryptedFunctions{
			DoCipherByReaderWriter: nil,
			CreateIV:               nil,
			CreateKey:              nil,
			CreatePermissionIV:     nil,
			ApplyPassphrase:        nil,
			DoMac:                  nil,
			DoRead:                 nil,
			NewByteRange:           nil,
			CreateRandomName:       createRandomName,
		}
		var actual = funcs.CreateRandomName()
		if actual != c.expected {
			t.Errorf("CreateRandomName())  %q, want %q", actual, c.expected)
		}
		fmt.Println(actual)
	}
}
