package utils

import (
	"faceit/domain/entity"
	"fmt"
	"strings"
)

func QueryBuilder(filter *entity.Filter, tableName string, page, pageSize int64) string {
	query := `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + tableName + ` WHERE`

	var conditions []string
	if filter.Country != "" {
		conditions = append(conditions, fmt.Sprintf("country = \"%s\"", filter.Country))
	}
	if !filter.CreatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("created_at = %s", filter.CreatedAt))
	}
	if !filter.UpdatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("updated_at = %s", filter.UpdatedAt))
	}

	joinedConditions := strings.Join(conditions, " AND ")

	query += " " + joinedConditions

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, page*pageSize)

	return query
}

func CountQueryBuilder(filter *entity.Filter, tableName string) string {
	query := `SELECT count(*) as total FROM ` + tableName + ` WHERE`

	var conditions []string
	if filter.Country != "" {
		conditions = append(conditions, fmt.Sprintf("country = \"%s\"", filter.Country))
	}
	if !filter.CreatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("created_at = %s", filter.CreatedAt))
	}
	if !filter.UpdatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("updated_at = %s", filter.UpdatedAt))
	}

	joinedConditions := strings.Join(conditions, " AND ")

	query += " " + joinedConditions

	return query
}

func UpdateQueryBuilder(user *entity.User, tableName string) string {
	query := `UPDATE ` + tableName + `SET `
	var updateFields []string
	if user.FirstName != "" {
		updateFields = append(updateFields, fmt.Sprintf("first_name = \"%s\"", user.FirstName))
	}
	if user.LastName != "" {
		updateFields = append(updateFields, fmt.Sprintf("last_name = \"%s\"", user.LastName))
	}
	if user.NickName != "" {
		updateFields = append(updateFields, fmt.Sprintf("nick_name = \"%s\"", user.NickName))
	}
	if user.Email != "" {
		updateFields = append(updateFields, fmt.Sprintf("email = \"%s\"", user.Email))
	}
	if user.Country != "" {
		updateFields = append(updateFields, fmt.Sprintf("country = \"%s\"", user.Country))
	}
	if user.Password != "" {
		updateFields = append(updateFields, fmt.Sprintf("password = \"%s\"", user.Password))
	}

	joinedUpdateFields := strings.Join(updateFields, " , ")
	query += " " + joinedUpdateFields + `, updated_at = NOW() ` + fmt.Sprintf("WHERE id = %d", user.ID)
	return query
}
