package config

type Config struct {
	App      App
	Database Database
}

type App struct {
	Host string
	Port int
}

type Database struct {
}
