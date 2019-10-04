package handler

import (
	"log"
	"net/http"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
)

func HandleTimeTable(writer http.ResponseWriter, request *http.Request) {
	timetables, err := service.Service.TimeTableService.TestTimeTable()
	if err != nil {
		log.Println(err)
		return
	}
	timetablesSize := len(timetables)
	if timetablesSize != 0 {
		log.Println(err)
	}
	rows := make([]TimeTableRow, timetablesSize)
	for i, j := range timetables {
		rows[i] = TimeTableRow{Num: j.Num, Name: j.Name}
	}
	response.Success(writer, TimeTableResponse{TimeTableRows: rows})

}

type TimeTableRow struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
}

type TimeTableResponse struct {
	TimeTableRows []TimeTableRow
}
