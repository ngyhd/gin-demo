package etc

type ServerConfig struct {
	//服务名称（唯一
	Name string `mapstructure:"name" json:"name"`
	//服务地址
	Host string `mapstructure:"host" json:"host"`
	//服务端口
	Port      int         `mapstructure:"port" json:"port"`
	RedisConf redisConfig `mapstructure:"redis" json:"redis"`
	MysqlConf mysqlConfig `mapstructure:"mysql" json:"mysql"`
	LogConf   logsConfig  `mapstructure:"logs" json:"logs"`
}

type redisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type mysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DB       string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type logsConfig struct {
	Path  string `mapstructure:"path" json:"path"`
	Level int    `mapstructure:"level" json:"level"`
	//最大保存时间（单位天
	MaxAge int `mapstructure:"max_age" json:"max_age"`
	//最大备份数
	MaxBackups int `mapstructure:"max_backups" json:"max_backups"`
	//最大Size MB
	MaxSize  int `mapstructure:"max_size" json:"max_size"`
	Compress int `mapstructure:"compress" json:"compress"`
}
