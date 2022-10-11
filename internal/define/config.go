package define

type Config struct {
	MaxIdleConns int `yaml:"maxIdleConns"`
	MaxOpenConns int `yaml:"maxOpenConns"`
}

func InitConfig() *Config {
	panic("not implement")
}
