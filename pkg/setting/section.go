package setting

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	MySQL struct {
		Host              string `yaml:"host"`
		Port              int    `yaml:"port"`
		User              string `yaml:"user"`
		Password          string `yaml:"password"`
		DBName            string `yaml:"dbname"`
		MaxIdleConns      int    `yaml:"maxIdleConns"`
		MaxOpenConns      int    `yaml:"maxOpenConns"`
		ConnexMaxLifetime int    `yaml:"connexMaxLifetime"`
	} `yaml:"mysql"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	Log LogConfig `yaml:"log"`
}

type LogConfig struct {
	LogLevel    string `yaml:"logLevel"`
	FileLogName string `yaml:"fileLogName"`
	MaxSize     int    `yaml:"maxSize"`
	MaxBackups  int    `yaml:"maxBackups"`
	MaxAge      int    `yaml:"maxAge"`
	Compress    bool   `yaml:"compress"`
}
