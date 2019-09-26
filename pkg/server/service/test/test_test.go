package test

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/test/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSceneTable := mock.NewMockSceneTable(ctrl)

	// 値が帰ってきたとき
	mockSceneTable.EXPECT().SelectRowsByName("test").Return(
		[]repository.SceneRow{
			repository.SceneRow{Name: "test", Num: 3, Action: "C", Text: "CCC"},
			repository.SceneRow{Name: "test", Num: 1, Action: "A", Text: "AAA"},
			repository.SceneRow{Name: "test", Num: 2, Action: "B", Text: "BBB"},
		}, nil)
	testService := TestService{
		SceneTable: mockSceneTable,
	}
	testValue, err := testService.Test("test")
	if err != nil {
		t.Fatal(err)
	}
	correctValue := []repository.SceneRow{
		repository.SceneRow{Name: "test", Num: 1, Action: "A", Text: "AAA"},
		repository.SceneRow{Name: "test", Num: 2, Action: "B", Text: "BBB"},
		repository.SceneRow{Name: "test", Num: 3, Action: "C", Text: "CCC"},
	}
	if !(reflect.DeepEqual(testValue, correctValue)) {
		t.Fatal("期待値:", correctValue, "出力:", testValue)
	}
}
