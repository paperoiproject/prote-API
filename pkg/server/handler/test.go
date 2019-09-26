package handler

import (
	"log"
	"net/http"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
)

// HandleTest /testハンドラの設定
func HandleTest(writer http.ResponseWriter, request *http.Request) {
	Form := TestRequest{
		Name: request.FormValue("name"),
	}
	scenes, err := service.Service.TestService.Test(Form.Name)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	scenesSize := len(scenes)
	if scenesSize == 0 {
		log.Println(err)
		response.BadRequest(writer, "不正なシナリオ名")
	}
	sceneRows := make([]SceneRow, scenesSize, scenesSize)
	for i, v := range scenes {
		sceneRows[i] = SceneRow{Name: v.Name, Num: v.Num, Action: v.Action, Text: v.Text}
	}
	response.Success(writer, TestResponse{SceneRows: sceneRows})
}

type SceneRow struct {
	Name   string `json:"name"`
	Num    int    `json:"num"`
	Action string `json:"action"`
	Text   string `json:"text"`
}

type TestRequest struct {
	Name string `json:"name"`
}

type TestResponse struct {
	SceneRows []SceneRow `json:"name"`
}
