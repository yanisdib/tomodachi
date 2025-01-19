package platform

var (
	createPlatformQuery = `INSERT INTO platforms(
		name, 
		server, 
		nitro_level, 
		video_resolution, 
		network_quality, 
		access_url
	) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`
)
