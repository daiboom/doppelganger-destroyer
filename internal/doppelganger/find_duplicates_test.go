package doppelganger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	// 테스트용 디렉토리 생성
	testDir := "./mock/1"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("테스트 디렉토리를 생성하는 중 오류 발생: %v", err)
	}

	// 테스트용 파일 생성
	files := []struct {
		name    string
		content string
	}{
		{"example1.txt", "Hello World"},
		{"example2.txt", "Hello World"},
		{"unique.txt", "Unique Content"},
	}

	for _, file := range files {
		filePath := filepath.Join(testDir, file.name)
		err := os.WriteFile(filePath, []byte(file.content), 0644)
		if err != nil {
			t.Fatalf("%s 파일을 생성하는 중 오류 발생: %v", file.name, err)
		}
	}

	defer func() {
		// 테스트가 끝난 후 디렉토리 정리
		err := os.RemoveAll(testDir)
		if err != nil {
			t.Fatalf("테스트 디렉토리를 정리하는 중 오류 발생: %v", err)
		}
	}()

	t.Run("중복 파일 탐지", func(t *testing.T) {
		err := FindDuplicates(testDir)
		if err != nil {
			t.Errorf("중복 파일을 탐지하는 중 오류 발생: %v", err)
		}
	})
}
