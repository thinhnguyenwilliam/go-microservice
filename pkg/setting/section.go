package setting

type Config struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql"`
	JWT    JWTConfig    `yaml:"jwt"`
	Log    LogConfig    `yaml:"log"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MySQLConfig struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	DBName            string `yaml:"dbname"`
	MaxIdleConns      int    `yaml:"maxIdleConns"`
	MaxOpenConns      int    `yaml:"maxOpenConns"`
	ConnexMaxLifetime int    `yaml:"connexMaxLifetime"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

type LogConfig struct {
	LogLevel    string `yaml:"logLevel"`
	FileLogName string `yaml:"fileLogName"`
	MaxSize     int    `yaml:"maxSize"` // in megabytes
	MaxBackups  int    `yaml:"maxBackups"`
	MaxAge      int    `yaml:"maxAge"` // in days
	Compress    bool   `yaml:"compress"`
}
