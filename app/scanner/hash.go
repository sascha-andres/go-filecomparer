package scanner

import (
	"encoding/hex"
	"io"
	"os"

	"crypto/sha256"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
)

func hash(fileToHash filedb.File) (result string, err error) {
	file, err := os.Open(fileToHash.RelativePath)
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
