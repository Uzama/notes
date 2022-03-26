package container

import (
	"database/sql"
	"notes/domain/interfaces"
)

type Containers struct {
	Adapters     Adapters
	Repositories Repositories
	Caches       Caches
}

type Adapters struct {
	Db *sql.DB
}

type Repositories struct {
	Note interfaces.NoteRepository
}

type Caches struct {
}
