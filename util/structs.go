// Package util provides
package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// генерируем hash сумму из файла
func GetChecksum(tmp multipart.File) string {

	hash := md5.New()
	if _, err := io.Copy(hash, tmp); err != nil {
		return ""
	}

	checksum := hex.EncodeToString(hash.Sum(nil))

	return checksum
}

// создаем директорию на сервере
func CreateDir(source string, dir string) (string, error) {
	dest := source + "/" + dir
	err := os.Mkdir(dest, 0755)
	if err != nil {
		return "", err
	}

	return dest, nil
}

// Get version & suffix from file name
func GetSuffix(appVersion string) bool {

	appVersion = appVersion[len(appVersion)-2:]

	suffix, _ := strconv.Atoi(appVersion)

	return suffix != 00
}
