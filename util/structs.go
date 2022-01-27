// Package util provides
package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"wemesse/conf"
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
func CreateDir(conf *conf.Conf, target string, ver string) (string, error) {
	dst := GenerateURI(conf.Deploy, target, ver)
	err := os.MkdirAll(dst, 0644)
	if err != nil {
		return "", err
	}

	return dst + "/", nil
}

// создаем директорию на сервере
func CreateFile(src string, name string, f multipart.File) error {
	dst, err := os.Create(src + name)
	if err != nil {

		return err
	}
	defer dst.Close()

	// копируем данные загруженного файла во вновь созданный файл в файловой системе
	if _, err := io.Copy(dst, f); err != nil {
		return err
	}

	return err
}

func GenerateURI(arr ...string) string {

	return strings.Join(arr, "/")
}

// Get version & suffix from file name
func GetSuffix(appVersion string) bool {

	appVersion = appVersion[len(appVersion)-2:]

	suffix, _ := strconv.Atoi(appVersion)

	return suffix != 00
}

// Преобразует размер в байтах в удобочитаемую строку
// в формате SI (десятичный)
func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
