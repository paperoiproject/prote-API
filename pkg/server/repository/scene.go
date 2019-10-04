package repository

import (
	"database/sql"
	"fmt"
	"log"
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

// BulkInsert データベースをレコードを登録する
func (scene *Scene) BulkInsert(name string, num int, works []string, texts []string) error {
	query := "INSERT INTO scene(name, num, action, text) VALUES"
	queryData := make([]interface{}, num*4, num*4)
	for i := 0; i < num*4; i = i + 4 {
		queryData[i] = name
		queryData[i+1] = i / 4
		queryData[i+2] = works[i/4]
		queryData[i+3] = texts[i/4]
		query += " (?, ?, ?, ?)"
		if i/4 == num-1 {
			break
		} else {
			query += ","
		}
	}
	log.Println(query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(queryData...)
	return err
}

// Delete データの削除
func (scene *Scene) Delete(name string) error {
	query := "DELETE FROM scene WHERE name = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(name)
	if err != nil {
		return err
	}
	checkNum, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if checkNum == 0 {
		return fmt.Errorf("消すデータが存在しません")
	}
	return err
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
