package timetable

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/timetable/mock"
	"sort"
)

type TimeTableService struct {
	TimeTable mock.TimeTable
}

func (timeTableService *TimeTableService) TestTimeTable() ([]repository.TimeTableRow, error) {
	timetable, err := timeTableService.TimeTable.SelectRow()
	if err != nil {
		return nil, err
	}
	sort.Slice(timetable, func(i, j int) bool {
		return timetable[i].Num < timetable[j].Num
	})
	return timetable, err
}

func (timeTableService *TimeTableService) Add(num []int, name []string) error {
	err := timeTableService.TimeTable.BulkInsert(num, name)
	return err
}

func (timeTableService *TimeTableService) Delete() error {
	err := timeTableService.TimeTable.Delete()
	return err
}
