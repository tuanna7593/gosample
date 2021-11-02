package config

import (
	"fmt"
	"time"
)

type Config struct {
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
}

type Server struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"` // second
}

type MySQL struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"db"`
	Port         string `yaml:"port"`
	MaxOpenConns int    `yaml:"max_open_cons"`
	MaxIdleConns int    `yaml:"max_idle_cons"`
}

// Conn return connection string
func (m *MySQL) Conn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&interpolateParams=true",
		m.User, m.Password, m.Host, m.Port, m.DB)
}
