package container

import (
	"database/sql"
	"notes/externals/adapters"
	"notes/externals/repositories"
	"notes/utils/config"
)

func Resolve(config config.Config) (Containers, error) {
	adaptrs, err := resolveAdapters(config)
	if err != nil {
		return Containers{}, err
	}

	repos, err := resolveRepostories(adaptrs.Db)
	if err != nil {
		return Containers{}, err
	}

	cont := Containers{
		Adapters:     adaptrs,
		Repositories: repos,
	}

	return cont, nil
}

func resolveAdapters(config config.Config) (Adapters, error) {

	mysql, err := adapters.NewDB(config.Database)
	if err != nil {
		return Adapters{}, err
	}

	adapters := Adapters{
		Db: mysql,
	}

	return adapters, nil
}

func resolveRepostories(db *sql.DB) (Repositories, error) {
	noteRepo := repositories.NewNoteRepository(db)

	repos := Repositories{
		Note: noteRepo,
	}

	return repos, nil
}
