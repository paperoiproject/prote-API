package handler

import (
	"fmt"
	"log"
	"net/http"
	"prote-API/pkg/server/service"
)

func HandleTimeTable(writer http.ResponseWriter, request *http.Request) {
	timetable, err := service.Service.TimeTableService.TestTimeTable()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, timetable)
}

type TimeTableRow struct {
	Num  int
	Name string
}

type TimeTableResponse struct {
	TimeTableRows []TimeTableRow
}
