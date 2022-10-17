package repository

const usersTableName = "users"

const (
	createUser = `INSERT INTO ` + usersTableName + ` SET first_name = ?, last_name = ?, nick_name = ?, password = ?, email = ?, country = ?`

	updateUser = `UPDATE ` + usersTableName + ` Set first_name = ?, last_name = ?, nick_name = ?, password = ?, email = ?, country = ?, updated_at = NOW() WHERE id = ?`

	deleteUser = `DELETE FROM ` + usersTableName + ` WHERE id = ?`

	getUsers = `SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM ` + usersTableName + ``
)
