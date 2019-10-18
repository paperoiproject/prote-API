package service

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/scene"
	"prote-API/pkg/server/service/test"
	"prote-API/pkg/server/service/timetable"
)

// service サービスの構造体
type service struct {
	TestService      *test.TestService
	SceneService     *scene.SceneService
	TimeTableService *timetable.TimeTableService
}

// Service サービスの生成(依存関係の解決)
var Service = service{
	TestService:      &test.TestService{SceneTable: &repository.Scene{}},
	SceneService:     &scene.SceneService{SceneTable: &repository.Scene{}},
	TimeTableService: &timetable.TimeTableService{TimeTable: &repository.TimeTable{}},
}
