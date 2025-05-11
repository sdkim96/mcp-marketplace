package config

type AppConfig struct {
	projectName   string
	projectSecret string
}

func (c *AppConfig) ProjectName() string {
	return c.projectName
}
func (c *AppConfig) ProjectSecret() string {
	return c.projectSecret
}

var Config = &AppConfig{
	projectName:   "MCP Marketplace",
	projectSecret: "secret",
}

func GetAppConfig() *AppConfig {

	// Must not be hardcoded in future
	return Config
}
