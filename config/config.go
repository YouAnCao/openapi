package config

// Config mapping
type Config struct {
	Server Server `mapstructure:"server"`
	Redis  Redis  `mapstructure:"redis"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

// Server config mapping
type Server struct {
	Port int
}

// Redis config mapping
type Redis struct {
	Host     string
	Port     int
	Password string
}

// Jwt 配置
type Jwt struct {
	SignKey string
}
