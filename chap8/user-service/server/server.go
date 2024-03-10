package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
	users "user-service/service"
)

type UserService struct {
	users.UnimplementedUsersServer
}

func (s *UserService) GetUser(ctx context.Context, in *users.UserGetRequest) (*users.UserGetReply, error) {
	// 수신 요청을 로그로 남기기
	log.Printf(
		"Received request for user with Email: %s Id: %s\n",
		in.Email,
		in.Id,
	)
	// 이메일 주소로부터 사용자 이름, 도메인 추출
	components := strings.Split(in.Email, "@")

	// 이메일 주소가 정상적이지 않은 경우 공백의 UserGetReply 값, 정상적으로 파싱하지 못하였다는 에러값을 반환
	if len(components) != 2 {
		return nil, errors.New("invalid email address")
	}

	// 더미 User 객체 생성
	u := users.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       36,
	}

	// UserGetReply 값, nil의 에러값을 반환
	return &users.UserGetReply{User: &u}, nil
}

func registerServices(s *grpc.Server) {
	users.RegisterUsersServer(s, &UserService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":50051"
	}

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	registerServices(s)
	log.Fatal(startServer(s, lis))
}
