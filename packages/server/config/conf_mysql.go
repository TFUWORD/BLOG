package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
  Port     int    `yaml:"port"`
  Db       string `yaml:"db"`
  User     string `yaml:"user"`
	Password string `yaml:"password"`
  LogLevel string `yaml:"log_level"`
}

func (m Mysql) Dsn() string{
  return "host=" + m.Host + " port=" + strconv.Itoa(m.Port) + " user=" + m.User + " dbname=" + m.Db + " password=" + m.Password + " sslmode=disable"
  // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  // return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
}
