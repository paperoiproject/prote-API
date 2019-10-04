package mock

import (
	"prote-API/pkg/server/repository"
)

// SceneTable SceneTableのmock
type SceneTable interface {
	SelectRowsByName(name string) ([]repository.SceneRow, error)
	BulkInsert(name string, num int, works []string, texts []string) error
	Delete(name string) error
}
