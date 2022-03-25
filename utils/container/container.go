package container

import "notes/domain/interfaces"

type Containers struct {
	Adapters     Adapters
	Repositories Repositories
	Caches       Caches
}

type Adapters struct {
}

type Repositories struct {
	Note interfaces.NoteRepository
}

type Caches struct {
}
