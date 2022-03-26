package repositories

import (
	"context"
	"database/sql"
	"log"
	"notes/domain/entities"
	"notes/domain/interfaces"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) interfaces.NoteRepository {
	repo := &NoteRepository{
		db: db,
	}

	return repo
}

func (repo NoteRepository) GetAllUnarchivedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {

	offset := (paginator.Page - 1) * paginator.Size

	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM note WHERE archived = 0 ORDER BY id ASC
		LIMIT ? OFFSET ?;
	`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, paginator.Size, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := make([]entities.Note, 0)

	for rows.Next() {
		note := entities.Note{}

		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Description,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (repo NoteRepository) GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {

	offset := (paginator.Page - 1) * paginator.Size

	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM note WHERE archived = 1 ORDER BY id ASC
		LIMIT ? OFFSET ?;
	`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, paginator.Size, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := make([]entities.Note, 0)

	for rows.Next() {
		note := entities.Note{}

		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Description,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (repo NoteRepository) GetSingle(ctx context.Context, id int64) (entities.Note, error) {

	note := entities.Note{}

	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM note WHERE id = ?;
	`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return note, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	err = row.Scan(
		&note.ID,
		&note.UserID,
		&note.Title,
		&note.Description,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return note, nil
	}

	if err != nil {
		return note, err
	}

	return note, nil
}

func (repo NoteRepository) CreateNote(ctx context.Context, note entities.Note) (int64, error) {

	query := `INSERT INTO note (user_id, title, description) VALUES (?, ?, ?);`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, note.UserID, note.Title, note.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo NoteRepository) ArchiveNote(ctx context.Context, id int64) (bool, error) {

	query := `UPDATE note SET archived = 1 where id = ?;`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo NoteRepository) UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error) {

	query := `
		UPDATE note SET title = ?, description = ?
		WHERE id = ?;
	`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newNote.Title, newNote.Description, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo NoteRepository) DeleteNote(ctx context.Context, id int64) (bool, error) {

	query := `DELETE FROM note WHERE id = ?;`

	stmt, err := repo.db.PrepareContext(ctx, query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo NoteRepository) IsExists(ctx context.Context, id int64) (bool, error) {

	query := `SELECT EXISTS(SELECT 1 FROM note WHERE id = ?;)`

	stmt, err := repo.db.PrepareContext(ctx, query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var exists bool

	err = row.Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (repo NoteRepository) IsArchived(ctx context.Context, id int64) (bool, error) {

	query := `SELECT EXISTS(SELECT 1 FROM note WHERE archived = 1;)`

	stmt, err := repo.db.PrepareContext(ctx, query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var exists bool

	err = row.Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return exists, nil
}
