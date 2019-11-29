package models

type Config struct {
	Addr string `yaml:"addr"`
	Mode string `yaml:"mode"`
	DB   DB     `yaml:"db"`
}

type DB struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DbName       string `yaml:"dbname"`
	Log          bool   `yaml:"log"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}
