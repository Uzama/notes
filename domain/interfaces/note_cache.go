package interfaces

import "notes/domain/entities"

type NoteCache interface {
	SetNote(note entities.Note)
	GetNote(id int64) (entities.Note, bool)
	DeleteNote(id int64)
	ArchiveNote(id int64)
	UnArchiveNote(id int64)
	IsArchived(id int64) bool
	Refresh()
}
