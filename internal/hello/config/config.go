package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`

	API     API     `yaml:"api"`
	Usecase Usecase `yaml:"usecase"`
	DB      DB      `yaml:"db"`
}

type API struct {
	MaxMessageSize int `yaml:"max_message_size"`
}

type Usecase struct {
	DefaultMessage string `yaml:"default_message"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}
