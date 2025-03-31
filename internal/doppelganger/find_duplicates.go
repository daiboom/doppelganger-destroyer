package doppelganger

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sync"
)

func FindDuplicates(rootDir string) error {
    fileMap := make(map[string][]string)
    nameMap := make(map[string][]string)
    var wg sync.WaitGroup
    var mu sync.Mutex
    var nameMu sync.Mutex

    err := filepath.WalkDir(rootDir, func(path string, entry os.DirEntry, err error) error {
        if err != nil {
            if os.IsPermission(err) {
                log.Printf("%s에 대한 권한이 거부되었습니다. 건너뜁니다...\n", path)
                if entry.IsDir() {
                    return filepath.SkipDir
                }
                return nil
            }
            return err
        }

        if !entry.IsDir() {
            wg.Add(1)
            go func(path string, entry os.DirEntry) {
                defer wg.Done()
                hash, err := GetFileHash(path) // 같은 패키지 내 함수 호출
                if err != nil {
                    log.Printf("%s의 해시를 계산하는 중 오류 발생: %v\n", path, err)
                    return
                }

                mu.Lock()
                fileMap[hash] = append(fileMap[hash], path)
                mu.Unlock()

                nameMu.Lock()
                nameMap[entry.Name()] = append(nameMap[entry.Name()], path)
                nameMu.Unlock()
            }(path, entry)
        }

        return nil
    })

    if err != nil {
        return err
    }

    wg.Wait()

    fmt.Println("해시 중복 파일 목록:")
    for hash, paths := range fileMap {
        if len(paths) > 1 {
            fmt.Printf("===============================================================\n")
            fmt.Printf("중복된 해시 %s이(가) 파일에서 발견되었습니다:\n", hash)
            fmt.Printf("===============================================================\n")
            
            for _, path := range paths {
                fmt.Println(path)
            }

            fmt.Printf("\n")
        }
    }

    fmt.Println("\v파일명 중복인 파일 목록:")
    for name, paths := range nameMap {
        if len(paths) > 1 {

            fmt.Printf("===============================================================\n")
            fmt.Printf("중복된 이름 %s이(가) 파일에서 발견되었습니다:\n", name)
            fmt.Printf("===============================================================\n")
            for _, path := range paths {
                fmt.Println(path)
            }
            fmt.Printf("\n")
        }
    }

    return nil
}
