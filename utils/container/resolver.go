package container

import (
	"database/sql"
	"notes/domain/interfaces"
	"notes/externals/adapters"
	"notes/externals/cache"
	"notes/externals/repositories"
	"notes/utils/config"
)

func Resolve(config config.Config) (Containers, error) {
	adaptrs, err := resolveAdapters(config)
	if err != nil {
		return Containers{}, err
	}

	caches := resolveCaches()

	repos, err := resolveRepostories(adaptrs.Db, caches.Note)
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

func resolveCaches() Caches {
	noteCache := cache.NewNoteCache()

	caches := Caches{
		Note: noteCache,
	}

	return caches
}

func resolveRepostories(db *sql.DB, noteCache interfaces.NoteCache) (Repositories, error) {
	noteRepo := repositories.NewNoteRepository(db, noteCache)

	repos := Repositories{
		Note: noteRepo,
	}

	return repos, nil
}
