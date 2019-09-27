package mock

import (
  "prote-API/pkg/server/repository"
)

type TimeTable interface{
  SelectRow() ([]repository.TimeTableRow, error)
}
