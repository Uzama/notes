package cache

import (
	"notes/domain/entities"
	"notes/domain/interfaces"
)

type NoteCache struct {
	noteIndex    map[int64]entities.Note
	archiveIndex map[int64]struct{}
}

func NewNoteCache() interfaces.NoteCache {

	cache := &NoteCache{
		noteIndex:    make(map[int64]entities.Note),
		archiveIndex: make(map[int64]struct{}),
	}

	return cache
}

func (cache *NoteCache) SetNote(note entities.Note) {
	cache.noteIndex[note.ID] = note
}

func (cache NoteCache) GetNote(id int64) (entities.Note, bool) {
	note, ok := cache.noteIndex[id]
	if !ok {
		return entities.Note{}, false
	}

	return note, true
}

func (cache *NoteCache) DeleteNote(id int64) {
	delete(cache.noteIndex, id)
}

func (cache *NoteCache) ArchiveNote(id int64) {
	cache.archiveIndex[id] = struct{}{}
}

func (cache *NoteCache) UnArchiveNote(id int64) {
	delete(cache.archiveIndex, id)
}

func (cache NoteCache) IsArchived(id int64) bool {
	_, ok := cache.archiveIndex[id]

	return ok
}

func (cache *NoteCache) Refresh() {
	cache.noteIndex = make(map[int64]entities.Note)
	cache.archiveIndex = make(map[int64]struct{})
}
