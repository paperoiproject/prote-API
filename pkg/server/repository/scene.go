package repository

import (
	"database/sql"
	"prote-API/pkg/server/repository/db"
)

// Scene Sceneテーブル
type Scene struct{}

// SelectRowsByName 名前が一致したシーンのSELECT
func (scene *Scene) SelectRowsByName(name string) ([]SceneRow, error) {
	rows, err := db.DB.Query("SELECT * FROM scene WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return convertRowsToSceneRows(rows)
}

// convertRowsToSceneRows rowsの[]SceneRowへの変換
func convertRowsToSceneRows(rows *sql.Rows) ([]SceneRow, error) {
	var sceneRows []SceneRow
	for rows.Next() {
		sceneRow := SceneRow{}
		err := rows.Scan(&sceneRow.Name, &sceneRow.Num, &sceneRow.Action, &sceneRow.Text)
		if err != nil {
			return nil, err
		}
		sceneRows = append(sceneRows, sceneRow)
	}
	return sceneRows, nil
}

// SceneRow sceneテーブルのrow全てを変換するのに使用
type SceneRow struct {
	Name   string
	Num    int
	Action string
	Text   string
}
