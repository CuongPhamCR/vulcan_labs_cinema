package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Log    LoggerSetting `mapstructure:"log"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LoggerSetting struct {
	LogLevel      string `mapstructure:"log_level"`
	FileLogName   string `mapstructure:"file_log_name"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxBackups    int    `mapstructure:"max_backups"`
	MaxAge        int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
	EnableConsole bool   `mapstructure:"enable_console"`
}
