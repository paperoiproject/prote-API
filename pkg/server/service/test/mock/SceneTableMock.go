package mock

import (
	"prote-API/pkg/server/repository"
)

// SceneTable SceneTableのmock
type SceneTable interface {
	SelectRowsByName(name string) ([]repository.SceneRow, error)
}
