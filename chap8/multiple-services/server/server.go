package main

import (
	"context"
	"errors"
	"log"
	svc "multiple-services/service"
	"strings"
)

type repoService struct {
	svc.UnimplementedServer
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

func (s *repoService) GetRepos(ctx context.Context, in *svc.RepoGetRequest) (*svc.RepoGetReply, error) {
	log.Printf(
		"Received request for repo with CreateId: %s Id: %s\n",
		in.CreatorId,
		in.Id,
	)

	repo := svc.Repository{
		Id:    in.Id,
		Name:  "text repo",
		Url:   "https://git.example.com/test/repo",
		Owner: &svc.User{Id: in.CreatorId, FirstName: "Jane"},
	}

	r := svc.RepoGetReply{
		Repo: []*svc.Repository{&repo},
	}
	return &r, nil
}

func register

func main() {

}
