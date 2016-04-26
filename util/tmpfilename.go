package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

const TEMP_FILE_NAME_RAND_LEN = 10

func TempFileName() (string, error) {
	b := make([]byte, TEMP_FILE_NAME_RAND_LEN)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	en := base64.StdEncoding
	d := make([]byte, en.EncodedLen(len(b)))
	en.Encode(d, b)
	return fmt.Sprintf("%s/%s", os.TempDir(), d), nil
}
