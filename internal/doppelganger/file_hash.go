package doppelganger

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func GetFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr // 파일 닫기 중 발생한 오류를 반환
		}
	}()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
