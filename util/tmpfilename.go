package util

import "io/ioutil"
import "os"

const TEMP_FILE_NAME_PREFIX = "gridfs"

func TempFileName() (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), TEMP_FILE_NAME_PREFIX)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return file.Name(), nil
}
