package models

type Config struct {
	Addr string `yaml:"addr"`
	Mode string `yaml:"mode"`
	DB   DB     `yaml:"db"`
	Log  Log    `yaml:"log"`
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

type Log struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}
