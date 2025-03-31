package main

import (
    "flag"
    "fmt"
    "log"

    "doppelganger-destroyer/internal/doppelganger"
)

func main() {
    // 중복 파일을 검색할 디렉토리를 지정하는 플래그
    rootDirPtr := flag.String("dir", ".", "중복 파일을 검색할 루트 디렉토리")
    helpPtr := flag.Bool("help", false, "도움말 메시지 표시")
    flag.Parse()

    // 도움말 플래그가 설정된 경우 사용법 출력
    if *helpPtr {
        fmt.Println("사용법:")
        fmt.Println("  -dir string: 중복 파일을 검색할 루트 디렉토리 (기본값: 현재 디렉토리)")
        fmt.Println("  -help: 이 도움말 메시지 표시")
        return
    }

    // 중복 파일 검색 실행
    if err := doppelganger.FindDuplicates(*rootDirPtr); err != nil {
        log.Fatal(err) // 오류 발생 시 프로그램 종료
    }
}
