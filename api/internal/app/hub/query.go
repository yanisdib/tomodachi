package hub

// CreateHubQuery is the SQL query to create a new hub
const CreateHubQuery = `
	INSERT INTO hubs (name, description, icon_url, tags, event_type, languages, user_limit, access_rule, password, start_at, end_at, active, open, created_at, server_id, host_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	RETURNING id, name, description, icon_url, tags, event_type, languages, user_limit, access_rule, start_at, end_at, active, open, created_at, server_id, host_id
	`

const GetHubsQuery = `
	SELECT id, name, description, icon_url, tags, event_type, languages, user_limit, access_rule, start_at, end_at, active, open, created_at, updated_at, closed_at, server_id, host_id
	FROM hubs
	ORDER BY created_at DESC
	`
