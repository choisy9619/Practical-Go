package main

import (
	"context"
	"errors"
	"log"
	"strings"
	users "user-service/service"
)

type UserService struct {
	users.UnimplementedUsersServer
}

func (s *UserService) GetUser(ctx context.Context, in *users.UserGetRequest) (*users.UserGetReply, error) {
	log.Printf(
		"Received request for user with Email: %s Id: %s\n",
		in.Email,
		in.Id,
	)
	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, errors.New("invalid email address")
	}
}

func main() {

}
