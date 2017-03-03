package scanner

import (
	"encoding/hex"
	"io"
	"os"

	"crypto/sha256"
)

func hash(fileInfo string) (result string, err error) {
	sugar.Debugw("hash", "fileInfo", fileInfo)

	file, err := os.Open(fileInfo)
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha256.New224()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	result = hex.EncodeToString(hash.Sum(nil))
	return
}
