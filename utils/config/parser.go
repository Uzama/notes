package config

func Parse() (Config, error) {
	app := &App{}
	db := &Database{}

	err := app.Parse()
	if err != nil {
		return Config{}, err
	}

	err = db.Parse()
	if err != nil {
		return Config{}, err
	}

	configs := Config{
		App:      *app,
		Database: *db,
	}

	return configs, nil
}
