package utils

import (
	entity2 "faceit/domain/user/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QueryBuilderTestSuite struct {
	suite.Suite
}

func (q *QueryBuilderTestSuite) TestQueryBuilder() {
	testCases := []struct {
		filter        *entity2.Filter
		tableName     string
		page          int64
		pageSize      int64
		expectedQuery string
	}{
		{
			filter: &entity2.Filter{
				Country: "UK",
			},
			tableName:     "users",
			page:          1,
			pageSize:      10,
			expectedQuery: "SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users WHERE country = \"UK\" ORDER BY id LIMIT 10 OFFSET 0",
		},
		{
			filter: &entity2.Filter{
				NickName: "test",
			},
			tableName:     "users",
			page:          1,
			pageSize:      10,
			expectedQuery: "SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users WHERE nick_name LIKE \"%test%\" ORDER BY id LIMIT 10 OFFSET 0",
		},
		{
			filter: &entity2.Filter{
				Country:  "UK",
				NickName: "test",
			},
			tableName:     "users",
			page:          1,
			pageSize:      10,
			expectedQuery: "SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users WHERE country = \"UK\" AND nick_name LIKE \"%test%\" ORDER BY id LIMIT 10 OFFSET 0",
		},
	}

	for _, tc := range testCases {
		query := QueryBuilder(tc.filter, tc.tableName, tc.page, tc.pageSize)
		assert.Equal(q.T(), tc.expectedQuery, query)
	}
}

func (q *QueryBuilderTestSuite) TestCountQueryBuilder() {
	testCases := []struct {
		filter        *entity2.Filter
		tableName     string
		expectedQuery string
	}{
		{
			filter: &entity2.Filter{
				Country: "UK",
			},
			tableName:     "users",
			expectedQuery: "SELECT count(*) as total FROM users WHERE country = \"UK\"",
		},
		{
			filter: &entity2.Filter{
				NickName: "test",
			},
			tableName:     "users",
			expectedQuery: "SELECT count(*) as total FROM users WHERE nick_name LIKE \"%test%\"",
		},
		{
			filter: &entity2.Filter{
				Country:  "UK",
				NickName: "test",
			},
			tableName:     "users",
			expectedQuery: "SELECT count(*) as total FROM users WHERE country = \"UK\" AND nick_name LIKE \"%test%\"",
		},
	}

	for _, tc := range testCases {
		query := CountQueryBuilder(tc.filter, tc.tableName)
		assert.Equal(q.T(), tc.expectedQuery, query)
	}
}

func (q *QueryBuilderTestSuite) TestUpdateQueryBuilder() {
	testCases := []struct {
		user          *entity2.User
		tableName     string
		expectedQuery string
	}{
		{
			user: &entity2.User{
				ID:        1,
				FirstName: "test",
			},
			tableName:     "users",
			expectedQuery: "UPDATE users SET first_name = \"test\" , updated_at = NOW() WHERE id = 1",
		},
		{
			user: &entity2.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
			},
			tableName:     "users",
			expectedQuery: "UPDATE users SET first_name = \"test\" , last_name = \"test\" , updated_at = NOW() WHERE id = 1",
		},
		{
			user: &entity2.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
			},
			tableName:     "users",
			expectedQuery: "UPDATE users SET first_name = \"test\" , last_name = \"test\" , nick_name = \"test\" , updated_at = NOW() WHERE id = 1",
		},
		{
			user: &entity2.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Country:   "UK",
			},
			tableName:     "users",
			expectedQuery: "UPDATE users SET first_name = \"test\" , last_name = \"test\" , nick_name = \"test\" , country = \"UK\" , updated_at = NOW() WHERE id = 1",
		},
	}

	for _, tc := range testCases {
		query := UpdateQueryBuilder(tc.user, tc.tableName)
		assert.Equal(q.T(), tc.expectedQuery, query)
	}
}

func TestQueryBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(QueryBuilderTestSuite))
}
