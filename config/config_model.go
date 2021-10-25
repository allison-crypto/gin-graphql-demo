package config

// Config for yaml file and environment variables
type Config struct {
	Server struct {
		Mode string
		Port int
	}
	Redis           redisConfig
	Mongo           mongoConfig
	TampManagerHost string
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

type mongoConfig struct {
	Addrs      []string
	Database   string
	ReplicaSet string
	Source     string
	Username   string
	Password   string
}
