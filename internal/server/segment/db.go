package segment

import (
	"database/sql"
	"fmt"
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
		return fmt.Errorf("error while creating a segment")
	}
	return nil
}

func (r *repository) Get(slug string) (Segment, error) {
	var segment Segment
	query := `SELECT id, slug FROM segments WHERE slug=$1`

	row := r.db.QueryRow(query, slug)
	err := row.Scan(&segment.Id, &segment.Slug)
	if err != nil {
		return Segment{}, fmt.Errorf("no segment with slug=%v", slug)
	}
	return segment, nil
}

func (r *repository) List() ([]Segment, error) {
	var segments []Segment
	query := `SELECT id, slug FROM segments`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while selecting segments")
	}

	for rows.Next() {
		var segment Segment
		err := rows.Scan(&segment.Id, &segment.Slug)
		if err != nil {
			return nil, fmt.Errorf("error while scanning segment")
		}
		segments = append(segments, segment)
	}
	return segments, nil
}

func (r *repository) Update(segment Segment) error {
	query := `UPDATE segments SET slug=$2 WHERE id=$1`

	_, err := r.db.Exec(query, segment.Id, segment.Slug)
	if err != nil {
		return fmt.Errorf("segment with id=%v not found or slug=%v isn`t correct",
			segment.Id, segment.Slug)
	}
	return nil
}

func (r *repository) Delete(id string) error {
	query := `DELETE FROM segments WHERE id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("user with id=%v not found", id)
	}
	return nil
}
