package doppelganger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileHash(t *testing.T) {
	// 테스트용 디렉토리 및 파일 생성
	testDir := "./mock/1"
	testFile := filepath.Join(testDir, "testfile.txt")

	// 테스트 파일 내용
	content := "Hello, World!"

	// 테스트 디렉토리 생성
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("테스트 디렉토리를 생성하는 중 오류 발생: %v", err)
	}

	// 테스트 파일 생성
	err = os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("테스트 파일을 생성하는 중 오류 발생: %v", err)
	}

	defer func() {
		// 테스트가 끝난 후 디렉토리 정리
		err := os.RemoveAll(testDir)
		if err != nil {
			t.Fatalf("테스트 디렉토리를 정리하는 중 오류 발생: %v", err)
		}
	}()

	t.Run("파일 해시 계산", func(t *testing.T) {
		hash, err := GetFileHash(testFile)
		if err != nil {
			t.Errorf("파일 해시를 계산하는 중 오류 발생: %v", err)
		}

		expectedHash := "65a8e27d8879283831b664bd8b7f0ad4" // "Hello, World!"의 MD5 해시값
		if hash != expectedHash {
			t.Errorf("해시값이 예상과 다릅니다. got: %s, want: %s", hash, expectedHash)
		}
	})

	t.Run("존재하지 않는 파일 처리", func(t *testing.T) {
		_, err := GetFileHash(filepath.Join(testDir, "nonexistent.txt"))
		if err == nil {
			t.Errorf("존재하지 않는 파일에 대해 오류가 발생하지 않았습니다.")
		}
	})
}
