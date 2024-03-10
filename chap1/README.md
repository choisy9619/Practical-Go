# CHAPTER 1. 커맨드 라인 애플리케이션 작성

## 1.1. 첫 애플리케이션
- 사용자 입력
- 입력값 검증
- 입력값 사용해 특정한 작업 수행
- 작업 수행 결과를 사용자에게 반환 (ex. 성공, 실패)


**getName() 함수**
- 첫 번째 매개변수 
  - **r** : io 패키지의 Reader 인터페이스를 구현하는 변수
  - 대개 프로그램을 실행하는 터미널 세션에서 프로그램의 표준 입력을 나타냄
  - os 패키지의 Stdin 변수
- 두 번째 매개변수
  - **w** : io 패키지의 Writer 인터페이스를 구현하는 변수
  - 대개 프로그램을 실행하는 터미널 세션에서 프로그램의 표준 출력을 나타냄
  - os 패키지의 Stdout 변수

<details><summary> 직접 Stdin, Stdout 변수를 참조하여 사용하지 않는 이유</summary>
unit 테스트에 불편해짐 - 입력을 특정하게 변경하는 것이 불가능 / 출력을 검증할 수 없기 때문 <br/>
함수의 매개변수로 인터페이스인 writer, reader를 주입해서 w,r로 제어할 수 있음
</details>    

- scanner : Scanner 타입의 변수 
  - Scan() 함수 : 개행 문자열을 읽을 때까지 데이터를 읽고 반환 
  - Text() 함수 : 읽은 데이터를 분자열로 반환

**parseArgs() 함수**
- 입력 매개변수 : 문자열의 슬라이스 []string
- 반환 : config 타입, error 타입
  -  config 구조체 : 메모리 내에 애플리케이션의 런타임 동작을 정의하는 데 사용

**validateArgs() 함수** - 입력값이 논리적으로 올바른지 확인

**runCmd() 함수** - 전체 로직 컨트롤 함수

**main() 함수** - os 패키지의 Exit() 함수 / 1의 종료 코드 exit code 로 호출 = 프로그램 종료


## 1.2 유닛 테스트 작성
- 표준 라이브러리의 testing 패키지 사용
- testConfig 구조체 정의
  - parse_args_test.go
  - validate_args_test.go
  - run_cmd_test.go 
  

**테스트 실행**
```
go test -v
```

**테스트 커버리지 확인**
```
go test -coverprofile cover.out
```

**테스트 커버리지 결과 html 확인**
```
go tool cover -html=cover.out
```


## 1.3 flag 패키지 사용

**일반적인 인터페이스 구조**
```
application [-h] [-n <value>] -silent <arg1> <arg2>
```
- -h : 안내 문구를 출력할지 지정하는 boolean 옵션값
- -n <value> : 사용자가 n이라는 옵션에 대해 지정하는 값
- -silent : -h 에 이어 또 다른 boolean 옵션값
- arg1, arg2 : 위치 인수 positional argument
  - 위치 인수 : 순서대로 인수를 전달하는 방법 - 순서가 달라지면 의미가 달라짐

**FlagSet 객체** : 커맨드 라인 애플리케이션의 인수를 처리하기 위한 추상 객체
  - FlagSet 객체의 메서드
    - Bool() : boolean 타입의 플래그를 정의
    - Int() : int 타입의 플래그를 정의
    - String() : string 타입의 플래그를 정의
    - Parse() : FlagSet 객체가 정의한 플래그를 파싱
    - Args() : FlagSet 객체가 정의한 플래그 이후의 위치 인수를 반환
    - NArg() : FlagSet 객체가 정의한 플래그 이후의 위치 인수의 개수를 반환
    - NFlag() : FlagSet 객체가 정의한 플래그의 개수를 반환


  - NewFlagSet() 함수
    - 첫 번째 매개변수 : 애플리케이션 자체의 이름
    - 두 번째 매개변수 : 커맨드 라인 인수를 파싱할 때 (fs.Parse() 함수 수행 중) 오류가 발생하는 경우 어떻게 처리할지 설정
      - ContinueOnError : Parse() 함수가 nil 외의 에러를 반환하더라도 프로그램을 계속 실행 - 파싱 에러를 직접 처리할 때 유용
      - ExitOnError : Parse() 함수가 nil 외의 에러를 반환하면 프로그램을 종료
      - PanicOnError : Parse() 함수가 nil 외의 에러를 반환하면 패닉 상태로 전환 
        - Panic의 경우, 프로그램이 종료되기 전에 recover() 함수를 사용해 마무리 정리 작업 cleanup action 수행 가능

    -   flag option 정의
        - init, float, string, bool, 커스텀 타입
        - 첫 번째 매개변수 : 해당 값을 저장할 변수의 주소값 (= 포인터)
        - 두 번째 매개변수 : 옵션 자체의 이름
        - 세 번째 매개변수 : 해당 옵션의 기본 값
        - 마지막 매개변수 : 사용자에게 보여줄 이 옵션의 목적을 문자열로 받음
          - 프로그램의 사용법을 출력할 때, 마지막 매개변수의 문자열이 자동으로 보이게 됨

**유닛 테스트**
- 함수 내에서 새로운 FlagSet 객체를 생성
- FlagSet 객체의 Output() 메소드를 사용해서 FlagSet의 모든 메서드가 지정된 io.Writer 객체의 w 변수로 출력
- 파싱할 인수를 매개변수 args로 전달


## 1.4 사용자 인터페이스 개선
- 중복된 오류 메시지를 제거
- 도움말 사용법 메시지를 사용자 정의
- 위치 인수를 통해 사용자의 이름을 입력 받음

