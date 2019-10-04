package repository

import (
  "database/sql"
	"prote-API/pkg/server/repository/db"
)

type TimeTable struct{}

type TimeTableRow struct{
  Num  int
  Name string
}

func (timeTable *TimeTable) SelectRow()([]TimeTableRow, error){
  rows, err := db.DB.Query("SELECT * FROM time_table;")
  if err != nil{
    return nil, err
  }
  return convertRowsToTimeTable(rows)
}

func convertRowsToTimeTable(rows *sql.Rows)([]TimeTableRow, error){
 var TimeTableRows []TimeTableRow
 for rows.Next(){
   timetable := TimeTableRow{}
   err := rows.Scan(&timetable.Num, &timetable.Name)
   if err != nil{
     return nil, err
   }
   TimeTableRows = append(TimeTableRows, timetable)
 }
 return TimeTableRows, nil
}
