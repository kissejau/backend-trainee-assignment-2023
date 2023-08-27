package user

import (
	"database/sql"
	"fmt"

	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/segment"
)

type repository struct {
	db *sql.DB
	sr segment.Repository
}

func NewRepository(db *sql.DB, sr segment.Repository) Repository {
	return &repository{
		db: db,
		sr: sr,
	}
}

func (r *repository) Create(name string) error {
	query := `INSERT INTO users (name) VALUES ($1)`

	_, err := r.db.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(id string) (User, error) {
	var user User
	query := `SELECT id, name FROM users WHERE id = $1`

	rows := r.db.QueryRow(query, id)
	if err := rows.Scan(&user.Id, &user.Name); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) List() ([]User, error) {
	var users []User
	query := `SELECT id, name FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) Update(u User) error {
	query := `UPDATE users SET name=$2 WHERE id=$1`

	_, err := r.db.Exec(query, u.Id, u.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id string) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}
	query := `DELETE FROM users WHERE id=$1`

	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetSegments(userId string) ([]segment.Segment, error) {
	var segments []segment.Segment
	query := `SELECT s.id, s.slug FROM users_segments AS us
	JOIN segments AS s ON s.id = us.segment_id
	JOIN users AS u ON u.id = us.user_id
	WHERE u.id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var segment segment.Segment

		err := rows.Scan(&segment.Id, &segment.Slug)
		if err != nil {
			return nil, err
		}

		segments = append(segments, segment)
	}
	return segments, nil
}

func (r *repository) SetSegments(setUserSegmentsDTO SetUserSegmentsDTO) error {
	var upSegments, downSegments []segment.Segment
	for _, slug := range setUserSegmentsDTO.UpSlugs {
		segment, err := r.sr.Get(slug)
		if err != nil {
			// log
			continue
		}

		upSegments = append(upSegments, segment)
	}

	for _, slug := range setUserSegmentsDTO.DownSlugs {
		segment, err := r.sr.Get(slug)
		if err != nil {
			// log
			continue
		}

		downSegments = append(downSegments, segment)
	}

	err := r.AddSegments(upSegments, setUserSegmentsDTO.UserId)
	if err != nil {
		return err
	}
	err = r.DeleteSegments(downSegments, setUserSegmentsDTO.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) AddSegments(segments []segment.Segment, userId string) error {
	for _, segment := range segments {
		err := r.AddSegment(segment.Id, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) AddSegment(segmentId, userId string) error {
	query := `INSERT INTO users_segments (user_id, segment_id) VALUES ($1, $2)`

	_, err := r.db.Exec(query, userId, segmentId)
	if err != nil {
		return fmt.Errorf("insert to users_segments failure")
	}
	return nil
}

func (r *repository) DeleteSegments(segments []segment.Segment, userId string) error {
	for _, segment := range segments {
		err := r.DeleteSegment(segment.Id, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) DeleteSegment(segmentId, userId string) error {
	query := `DELETE FROM users_segments AS us WHERE us.user_id=$1 AND us.segment_id=$2`

	_, err := r.db.Exec(query, userId, segmentId)
	if err != nil {
		return fmt.Errorf("delete from users_segments error")
	}
	return nil
}
