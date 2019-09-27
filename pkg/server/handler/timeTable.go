package handler

import (
	"log"
	"net/http"
	"prote-API/pkg/server/service"
	"prote-API/pkg/server/handler/response"

)

func HandleTimeTable(writer http.ResponseWriter, request *http.Request){
	timetable, err := service.Service.TimeTableService.TestTimeTable()
	if err != nil{
		log.Println(err)
	}
}

type TimeTableRow struct{
	Num  int
	Name string
}

type TimeTableResponse struct{
	TimeTableRows []TimeTableRow
}
