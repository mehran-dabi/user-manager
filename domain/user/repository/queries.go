package repository

const usersTableName = "users"

const (
	createUser = `INSERT INTO ` + usersTableName + ` SET first_name = ?, last_name = ?, nick_name = ?, password = ?, email = ?, country = ?`

	deleteUser = `DELETE FROM ` + usersTableName + ` WHERE id = ?`

	getUserByID = `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + usersTableName + ` WHERE id = ?`

	getUserByEmail = `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + usersTableName + ` WHERE email = ?`

	getUserByNickName = `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + usersTableName + ` WHERE nick_name = ?`
)
