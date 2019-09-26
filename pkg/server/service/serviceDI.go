package service

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/test"
)

type service struct {
	TestService *test.TestService
}

var Service = service{
	TestService: &test.TestService{SceneTable: &repository.Scene{}},
}
