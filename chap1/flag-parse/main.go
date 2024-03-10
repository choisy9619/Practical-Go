package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int // 출력 인사 횟수

}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h | --help]
A greeter application which prints the name you entered <integer> number of times.`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	// flagSet 객체 생성
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	// 첫 번째 flag 옵션 정의
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	// NArg() 메소드 : 플래그 옵션이 파싱된 이후에 주어진 위치 인수의 개수를 반환
	// 이 프로그램은 별도의 위치 인수를 필요로 하지 않으므로, 하나 이상의 값이 지정된 경우 에러를 출
	if fs.NArg() != 0 {
		return c, errors.New("Positional arguments specified")
	}
	return c, nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r) // 스캐너 타입의 변수를 생성

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()

	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}

	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nive to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {

	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
