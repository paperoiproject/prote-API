package test

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/test/mock"
	"sort"
)

type TestService struct {
	SceneTable mock.SceneTable
}

func (testService *TestService) Test(name string) ([]repository.SceneRow, error) {
	scenes, err := testService.SceneTable.SelectRowsByName(name)
	if err != nil {
		return nil, err
	}
	sort.Slice(scenes, func(i, j int) bool {
		return scenes[i].Num < scenes[j].Num
	})
	return scenes, err
}
