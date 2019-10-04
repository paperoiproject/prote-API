package test

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/test/mock"
	"sort"
)

type TimeTableService struct{
  TimeTable mock.TimeTable
}

func (timeTableService *TimeTableService) TestTimeTable()([]repository.TimeTableRow, error){
  timetable, err := timeTableService.TimeTable.SelectRow()
  if err != nil{
    return nil, err
  }
  sort.Slice(timetable, func(i, j int) bool {
		return timetable[i].Num < timetable[j].Num
	})
	return timetable, err
}
