package utils

const (
	// REGISTER USER
	REGISTER_USER = "INSERT INTO users(name, occoupation, email, password_hash, role, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6, $7)"
	FIND_BY_EMAIL = "SELECT id, name, email, password_hash FROM users WHERE email = $1"
	GET_USER_BY_ID = "SELECT id, name, email from users WHERE id = $1"
	UPLOAD_AVATAR = "INSERT INTO users avatar_filename "

	// CAMPAIGN
	FIND_CAMPAIGN_ALL = "SELECT ci.campaign_id, ci.file_name, ci.is_primary ,c.id, c.user_id, c.name, c.short_description, c.goal_amount, c.current_amount, c.slug FROM campaign AS c JOIN campaign_images AS ci ON c.id = ci.campaign_id"

	FIND_CAMPAIGN_BY_ID = "SELECT ci.campaign_id, ci.file_name, ci.is_primary , c.id, c.user_id, c.name, c.short_description, c.goal_amount, c.current_amoun, c.slug FROM campaign AS c JOIN campaign_images AS ci ON c.id = ci.campaign_id WHERE c.id = $1"
)