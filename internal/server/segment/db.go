package segment

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(segment Segment) error {
	query := `INSERT INTO segments (slug) VALUES ($1)`

	_, err := r.db.Exec(query, segment.Slug)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(slug string) (Segment, error) {
	var segment Segment
	query := `SELECT id, slug FROM segments WHERE slug=$1`

	row := r.db.QueryRow(query, slug)
	err := row.Scan(&segment.Id, &segment.Slug)
	if err != nil {
		return Segment{}, err
	}
	return segment, nil
}

func (r *repository) List() ([]Segment, error) {
	var segments []Segment
	query := `SELECT id, slug FROM segments`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var segment Segment
		err := rows.Scan(&segment.Id, &segment.Slug)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}
	return segments, nil
}

func (r *repository) Update(segment Segment) error {
	query := `UPDATE segments SET slug=$2 WHERE id=$1`

	_, err := r.db.Exec(query, segment.Id, segment.Slug)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id string) error {
	query := `DELETE FROM segments WHERE id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
