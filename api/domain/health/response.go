package health

type HealthCheckResponse struct {
	DatabaseStatus       string `json:"database_status"`
	ServerStatus         string `json:"server_status"`
	APIServerStatus      string `json:"api_server_status"`
	APIServerVersion     string `json:"api_server_version"`
	APIServerUptime      string `json:"api_server_uptime"`
	APIServerStartTime   string `json:"api_server_start_time"`
	APIServerHost        string `json:"api_server_host"`
	APIServerEnvironment string `json:"api_server_environment"`
	APIServerGoVersion   string `json:"api_server_go_version"`
}
