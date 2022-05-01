package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

const DEFAULT_CONFIG = `[General]
Port=5023

[SSL]
Use=false
Pem=./cart/server.pem
Key=./cart/server.key

[Database]
Host=127.0.0.1
Port=3306
User=root
Password=root
Database=schomem

[Service]
Register=true
`

type Config struct {
	General  *GeneralConfig
	SSL      *SSLConfig
	Database *DatabaseConfig
	Service  *ServiceConfig
}

type GeneralConfig struct {
	Port int
}

type SSLConfig struct {
	Use bool
	Pem string
	Key string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type ServiceConfig struct {
	Register bool
}

func NewConfigFile(path, value string) {
	_, err := os.Stat(path)

	if err != nil {
		// 文件不存在
		f, err := os.Create(path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			_, err = f.Write([]byte(value))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer f.Close()
		}
	}
}

func ParseConfig(path string) *Config {
	NewConfigFile(path, DEFAULT_CONFIG)
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := &Config{
		General: &GeneralConfig{
			Port: cfg.Section("General").Key("Port").MustInt(5023),
		},
		SSL: &SSLConfig{
			Use: cfg.Section("SSL").Key("Use").MustBool(false),
			Pem: cfg.Section("SSL").Key("Pem").MustString("./cart/server.pem"),
			Key: cfg.Section("SSL").Key("Key").MustString("./cart/server.key"),
		},
		Database: &DatabaseConfig{
			Host:     cfg.Section("Database").Key("Host").MustString("127.0.0.1"),
			Port:     cfg.Section("Database").Key("Port").MustInt(3306),
			User:     cfg.Section("Database").Key("User").MustString("root"),
			Password: cfg.Section("Database").Key("Password").MustString("root"),
			Database: cfg.Section("Database").Key("Database").MustString("schomem"),
		},
		Service: &ServiceConfig{
			Register: cfg.Section("Service").Key("Register").MustBool(true),
		},
	}
	return config
}
