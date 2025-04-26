package setting

type DBConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connexMaxLifetime"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	MySQL  DBConfig     `yaml:"mysql"`
	JWT    JWTConfig    `yaml:"jwt"`
}
