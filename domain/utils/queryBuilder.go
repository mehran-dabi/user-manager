package utils

import (
	"faceit/domain/entity"
	"fmt"
	"strings"
)

func QueryBuilder(filter *entity.Filter, tableName string, page, pageSize int64) string {
	query := `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at, count(*) over() as total_count FROM ` + tableName

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
