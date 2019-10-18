package handler

import (
	"fmt"
	"log"
	"net/http"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
	"strconv"
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

func PostHandleTimeTableAdd(writer http.ResponseWriter, request *http.Request) {
	FormSize, err := strconv.Atoi(request.FormValue("size"))
	if err != nil {
		log.Println(err)
		return
	}
	FormSize /= 2
	var Nums = make([]int, FormSize, FormSize)
	var Names = make([]string, FormSize, FormSize)

	for i := 0; i < FormSize; i++ {
		num := request.FormValue(fmt.Sprintf("num%v", i+1))
		name := request.FormValue(fmt.Sprintf("name%v", i+1))

		Nums[i], _ = strconv.Atoi(num)
		Names[i] = name
	}
	log.Println(Nums)
	log.Println(Names)

	err = service.Service.TimeTableService.Add(Nums, Names)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	response.Success(writer, ResponseTimeTableAdd{Nums: Nums, Names: Names})
}

type ResponseTimeTableAdd struct {
	Nums  []int    `json:"nums"`
	Names []string `json:"names"`
}

func PostHandleTimeTableDelete(writer http.ResponseWriter, request *http.Request) {
	err := service.Service.TimeTableService.Delete()
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}

	response.Success(writer, "OK")
}
