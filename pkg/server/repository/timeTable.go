package repository

import (
	"database/sql"
	"prote-API/pkg/server/repository/db"
)

type TimeTable struct{}

type TimeTableRow struct {
	Num  int
	Name string
}

func (timeTable *TimeTable) SelectRow() ([]TimeTableRow, error) {
	rows, err := db.DB.Query("SELECT * FROM time_table;")
	if err != nil {
		return nil, err
	}
	return convertRowsToTimeTable(rows)
}

func convertRowsToTimeTable(rows *sql.Rows) ([]TimeTableRow, error) {
	var TimeTableRows []TimeTableRow
	for rows.Next() {
		timetable := TimeTableRow{}
		err := rows.Scan(&timetable.Num, &timetable.Name)
		if err != nil {
			return nil, err
		}
		TimeTableRows = append(TimeTableRows, timetable)
	}
	return TimeTableRows, nil
}

func (timeTable *TimeTable) BulkInsert(num []int, name []string) error {
	length := len(num) * 2
	query := "INSERT INTO time_table(num, name) VALUES"
	queryData := make([]interface{}, length, length)
	for i := 0; i < length; i += 2 {
		queryData[i] = num[i/2]
		queryData[i+1] = name[i/2]
		query += "(?, ?)"
		if i/2 == len(num)-1 {
			break
		} else {
			query += ","
		}
	}

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(queryData...)
	return err
}

func (timeTable *TimeTable) Delete() error {
	_, err := db.DB.Exec("TRUNCATE TABLE time_table;")
	return err
}
