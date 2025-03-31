# Doppelganger Destroyer

중복 파일을 탐지하고 제거하는 Go 기반 프로그램입니다. 이 도구는 파일의 해시값과 이름을 기준으로 중복된 파일을 식별합니다.

---

## 프로젝트 구조
```
project-root/
├── cmd/ # 실행 파일 관련 코드
│ ├── doppelganger-destroyer/
│ ├── main.go # 프로그램의 진입점
│ └── find_duplicates.go # 중복 파일 탐지 로직
├── internal/ # 내부 패키지 (외부에서 접근 불가)
│ ├── doppelganger/
│ ├── file_hash.go # 파일 해시 계산 로직
│ └── utils/ # 유틸리티 관련 코드
├── mock/ # 테스트를 위한 모의 데이터
│ ├── example.txt # 테스트용 파일
│ └── example2.txt # 추가 테스트용 파일
├── .gitignore # Git에서 제외할 파일 목록
├── go.mod # Go 모듈 설정
├── Makefile # 빌드 및 실행 명령어 정의
└── README.md # 프로젝트 설명 파일
```


---

## 주요 기능

1. **중복 탐지**:
   - 파일의 MD5 해시값을 계산하여 중복된 파일을 식별합니다.
   - 동일한 이름을 가진 파일도 탐지합니다.

2. **효율적인 병렬 처리**:
   - Goroutine과 `sync` 패키지를 사용하여 빠른 탐지를 지원합니다.

---

## 설치 및 실행 방법

1. **프로젝트 클론**
```bash
git clone https://github.com/daiboom/doppelganger-destroyer.git
cd doppelganger-destroyer
```
2. **종속성 설치**
```bash
go mod tidy
```

3. **프로그램 실행**
```bash
go run cmd/doppelganger-destroyer/main.go -dir /path/to/your/directory
```

4. **Makefile 사용**
- 실행:
  ```
  make run DIR=/path/to/your/directory
  ```
- 린트:
  ```bash
   # macOS 
   brew install golangci-lint
  ```
  ```bash
  make lint
  ```
- 테스트:

  ```
  make test
  ```
