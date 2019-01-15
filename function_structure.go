package fp

import (
	"io"

	"go.uber.org/zap"
)

//This is a structure from crypto.go, not a function
type ByteRange struct {
	Start int64
	Stop  int64
}

type DoCipherByReaderWriter func(
	logger *zap.Logger,
	inFile io.Reader,
	outFile io.Writer,
	key []byte,
	iv []byte,
	description string,
	byteRange *ByteRange,
) (checksum []byte, length int64, err error)

type CreateIV func() (iv []byte)

type CreateKey func() (key []byte)

type CreatePermissionIV func() (key []byte)

type ApplyPassphrase func(passphrase string, permissionIV, fileKey []byte) []byte

type DoMac func(passphrase string, permissionIV []byte, grantee string, c, r, u, d, s bool, encryptedKey []byte) []byte

//TODO See if you can figure out how to make this a method with the *cipherStreamReader.  Won't compile with that extra curried param.
type DoRead func(dst []byte) (n int, err error)

type NewByteRange func() *ByteRange

type CreateRandomName func() string

type EncryptedFunctions struct {
	DoCipherByReaderWriter DoCipherByReaderWriter
	CreateIV               CreateIV
	CreateKey              CreateKey
	CreatePermissionIV     CreatePermissionIV
	ApplyPassphrase        ApplyPassphrase
	DoMac                  DoMac
	DoRead                 DoRead
	NewByteRange           NewByteRange
	CreateRandomName       CreateRandomName
}

func createRandomName() string {
	return "fred"
}
