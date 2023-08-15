package utils

const (
	// REGISTER USER
	REGISTER_USER = "INSERT INTO users(name, accoupation, email, password_hash, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)"
	FIND_BY_EMAIL = "SELECT id, name, email FROM users WHERE email = $1"
	GET_USER_BY_ID = "SELECT id from users WHERE id = $1"
)