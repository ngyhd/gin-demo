package etc

type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`   // 服务名称
	Host      string      `mapstructure:"host" json:"host"`   // 主机地址
	Port      int         `mapstructure:"port" json:"port"`   // 启动端口
	RedisConf redisConfig `mapstructure:"redis" json:"redis"` // Redis配置
	MysqlConf mysqlConfig `mapstructure:"mysql" json:"mysql"` // Mysql配置
	LogConf   logsConfig  `mapstructure:"logs" json:"logs"`   // 日志配置
}

type redisConfig struct {
	Host     string `mapstructure:"host" json:"host"`         // Redis地址。集群用多个逗号分割
	Port     string `mapstructure:"port" json:"port"`         // Redis端口
	Password string `mapstructure:"password" json:"password"` // Redis密码
}

type mysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`         // Mysql地址
	Port     int    `mapstructure:"port" json:"port"`         // Mysql端口
	DB       string `mapstructure:"db" json:"db"`             // 数据库
	User     string `mapstructure:"user" json:"user"`         // Mysql用户
	Password string `mapstructure:"password" json:"password"` // Mysql密码
}

type logsConfig struct {
	Path       string `mapstructure:"path" json:"path"`               // 配置文件路径
	Level      int    `mapstructure:"level" json:"level"`             // 日志级别
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`         // 最大保存时间（单位天
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"` //最大备份数
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`       // 最大Size MB
	Compress   int    `mapstructure:"compress" json:"compress"`       // 是否压缩
}
