package utils

import (
	entity2 "faceit/domain/user/entity"
	"fmt"
	"strings"
)

func QueryBuilder(filter *entity2.Filter, tableName string, page, pageSize int64) string {
	query := `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + tableName

	var conditions []string
	if filter.Country != "" {
		conditions = append(conditions, fmt.Sprintf("country = \"%s\"", filter.Country))
	}
	if filter.NickName != "" {
		conditions = append(conditions, fmt.Sprintf("nick_name LIKE \"%%%s%%\"", filter.NickName))
	}

	joinedConditions := strings.Join(conditions, " AND ")
	if joinedConditions != "" {
		query += " WHERE " + joinedConditions
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT %d OFFSET %d", pageSize, (page-1)*pageSize)

	return query
}

func CountQueryBuilder(filter *entity2.Filter, tableName string) string {
	query := `SELECT count(*) as total FROM ` + tableName

	var conditions []string
	if filter.Country != "" {
		conditions = append(conditions, fmt.Sprintf("country = \"%s\"", filter.Country))
	}
	if filter.NickName != "" {
		conditions = append(conditions, fmt.Sprintf("nick_name LIKE \"%%%s%%\"", filter.NickName))
	}

	joinedConditions := strings.Join(conditions, " AND ")
	if joinedConditions != "" {
		query += " WHERE " + joinedConditions
	}

	return query
}

func UpdateQueryBuilder(user *entity2.User, tableName string) string {
	query := `UPDATE ` + tableName + ` SET`
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
	query += " " + joinedUpdateFields + ` , updated_at = NOW() ` + fmt.Sprintf("WHERE id = %d", user.ID)
	return query
}
