package models

import (
	"fmt"
	"main/db"
	"main/types"
	"strings"
	"time"
)

type UserModel interface {
	GetUsers(types.ReqUser) ([]types.User, error)
	GetUserByID(int) (types.User, error)
	CreateUser(types.ReqCreateUser) (int, error)
}

type userModel struct {
	db db.Database
}

func NewUserModel(db db.Database) UserModel {
	return &userModel{db}
}

func (um *userModel) GetUsers(req types.ReqUser) ([]types.User, error) {
	results := []types.User{}
	db := um.db.GetDB()
	query := `
		SELECT
			id,
			name,
			email,
			created_at,
			updated_at
		FROM	
			users
	`

	if req.Order != "" {
		switch req.Order {
		case "id":
			query += fmt.Sprintf(" ORDER BY %s ", "id")
		case "created_at":
			query += fmt.Sprintf(" ORDER BY %s ", "created_at DESC")
		default:
			query += fmt.Sprintf(" ORDER BY %s ", "id")
		}
	}

	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d ", req.Limit)
	}

	rows, err := db.Query(query)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		result := types.User{}
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, err
}

func (um *userModel) GetUserByID(id int) (types.User, error) {
	result := types.User{}
	db := um.db.GetDB()
	query := `
		SELECT
			id,
			name,
			email,
			created_at,
			updated_at
		FROM	
			users
		WHERE
			id = ?
	`

	rows, err := db.Query(query, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return result, err
		}
	}

	return result, err
}

func (um *userModel) CreateUser(req types.ReqCreateUser) (int, error) {
	id := 0
	db := um.db.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return id, err
	}

	query := `
		INSERT INTO 
			users (
				name,
				email,
				created_at
			) VALUES 
	`
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	fmtNowJST := time.Now().In(jst).Format(time.RFC3339)

	value := []interface{}{
		req.Name,
		req.Email,
		fmtNowJST,
	}

	query += um.generatePlaceholder(len(value))

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(value...)
	if err != nil {
		tx.Rollback()
		return id, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return id, err
	}

	lastInsertId, err := result.LastInsertId()
	return int(lastInsertId), err
}

func (um *userModel) generatePlaceholder(fieldCount int) string {
	tmp := strings.Repeat("?,", fieldCount)
	tmp = strings.TrimRight(tmp, ",")
	return fmt.Sprintf("(%s)", tmp)
}
