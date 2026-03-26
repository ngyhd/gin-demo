package config

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

func TestServerConfigStruct(t *testing.T) {
	cfg := ServerConfig{
		Name: "test-service",
		Host: "localhost",
		Port: 8080,
		Mode: "debug",
		RedisConf: redisConfig{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
		},
		MysqlConf: mysqlConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "test_db",
			User:     "root",
			Password: "password",
		},
		LogConf: logsConfig{
			Path:       "./logs",
			Level:      "debug",
			MaxAge:     7,
			MaxBackups: 7,
			MaxSize:    100,
			Compress:   true,
		},
	}

	if cfg.Name != "test-service" {
		t.Errorf("Name = %v, want test-service", cfg.Name)
	}
	if cfg.Host != "localhost" {
		t.Errorf("Host = %v, want localhost", cfg.Host)
	}
	if cfg.Port != 8080 {
		t.Errorf("Port = %v, want 8080", cfg.Port)
	}
	if cfg.Mode != "debug" {
		t.Errorf("Mode = %v, want debug", cfg.Mode)
	}
}

func TestRedisConfigStruct(t *testing.T) {
	rc := redisConfig{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "secret",
	}

	if rc.Host != "127.0.0.1" {
		t.Errorf("RedisConfig.Host = %v, want 127.0.0.1", rc.Host)
	}
	if rc.Port != "6379" {
		t.Errorf("RedisConfig.Port = %v, want 6379", rc.Port)
	}
	if rc.Password != "secret" {
		t.Errorf("RedisConfig.Password = %v, want secret", rc.Password)
	}
}

func TestMysqlConfigStruct(t *testing.T) {
	mc := mysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "gin_demo",
		User:     "root",
		Password: "123456",
	}

	if mc.Host != "127.0.0.1" {
		t.Errorf("MysqlConfig.Host = %v, want 127.0.0.1", mc.Host)
	}
	if mc.Port != 3306 {
		t.Errorf("MysqlConfig.Port = %v, want 3306", mc.Port)
	}
	if mc.DB != "gin_demo" {
		t.Errorf("MysqlConfig.DB = %v, want gin_demo", mc.DB)
	}
	if mc.User != "root" {
		t.Errorf("MysqlConfig.User = %v, want root", mc.User)
	}
	if mc.Password != "123456" {
		t.Errorf("MysqlConfig.Password = %v, want 123456", mc.Password)
	}
}

func TestLogsConfigStruct(t *testing.T) {
	lc := logsConfig{
		Path:       "./logs",
		Level:      "debug",
		MaxAge:     7,
		MaxBackups: 7,
		MaxSize:    100,
		Compress:   true,
	}

	if lc.Path != "./logs" {
		t.Errorf("LogsConfig.Path = %v, want ./logs", lc.Path)
	}
	if lc.Level != "debug" {
		t.Errorf("LogsConfig.Level = %v, want debug", lc.Level)
	}
	if lc.MaxAge != 7 {
		t.Errorf("LogsConfig.MaxAge = %v, want 7", lc.MaxAge)
	}
	if lc.MaxBackups != 7 {
		t.Errorf("LogsConfig.MaxBackups = %v, want 7", lc.MaxBackups)
	}
	if lc.MaxSize != 100 {
		t.Errorf("LogsConfig.MaxSize = %v, want 100", lc.MaxSize)
	}
	if !lc.Compress {
		t.Errorf("LogsConfig.Compress = %v, want true", lc.Compress)
	}
}

func TestViperConfigLoading(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")

	yamlContent := `
name: test-app
host: 0.0.0.0
port: 9090
mode: debug
redis:
  host: localhost
  port: "6380"
  password: redis-pass
mysql:
  host: localhost
  port: 3307
  db: test_db
  user: testuser
  password: testpass
logs:
  path: /var/log
  level: info
  max_age: 14
  max_backups: 14
  max_size: 500
  compress: false
`

	err := v.ReadConfig(bytes.NewReader([]byte(yamlContent)))
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}

	var cfg ServerConfig
	err = v.Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	if cfg.Name != "test-app" {
		t.Errorf("Name = %v, want test-app", cfg.Name)
	}
	if cfg.Port != 9090 {
		t.Errorf("Port = %v, want 9090", cfg.Port)
	}
	if cfg.RedisConf.Host != "localhost" {
		t.Errorf("RedisConf.Host = %v, want localhost", cfg.RedisConf.Host)
	}
	if cfg.RedisConf.Port != "6380" {
		t.Errorf("RedisConf.Port = %v, want 6380", cfg.RedisConf.Port)
	}
	if cfg.MysqlConf.DB != "test_db" {
		t.Errorf("MysqlConf.DB = %v, want test_db", cfg.MysqlConf.DB)
	}
	if cfg.LogConf.Level != "info" {
		t.Errorf("LogConf.Level = %v, want info", cfg.LogConf.Level)
	}
	if cfg.LogConf.Compress != false {
		t.Errorf("LogConf.Compress = %v, want false", cfg.LogConf.Compress)
	}
}

