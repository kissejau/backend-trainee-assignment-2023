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

func (r *repository) GetSegments(userId string) ([]SegmentWithActiveStatusDTO, error) {
	var segments []SegmentWithActiveStatusDTO
	query := `SELECT s.id, s.slug, us.deadline FROM users_segments AS us
	JOIN segments AS s ON s.id = us.segment_id
	JOIN users AS u ON u.id = us.user_id
	WHERE u.id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var segment SqlableSegmentDTO
		err := rows.Scan(&segment.Id, &segment.Slug, &segment.Deadline)

		if err != nil {
			return nil, err
		}

		segments = append(segments, *segment.SegmentWithActiveStatusDTO())
	}
	return segments, nil
}

func (r *repository) SetSegments(setUserSegmentsDTO SqlableSetUserSegmentsDTO) error {
	var upSegments, downSegments []SqlableSegmentDTO
	for _, slug := range setUserSegmentsDTO.UpSlugs {
		segment, err := r.sr.Get(slug.Slug)
		if err != nil {
			// log
			continue
		}

		upSegments = append(upSegments, SqlableSegmentDTO{
			Id:       segment.Id,
			Slug:     segment.Slug,
			Deadline: slug.Deadline,
		})
	}

	for _, slug := range setUserSegmentsDTO.DownSlugs {
		segment, err := r.sr.Get(slug.Slug)
		if err != nil {
			// log
			continue
		}

		downSegments = append(downSegments, SqlableSegmentDTO{
			Id:       segment.Id,
			Slug:     segment.Slug,
			Deadline: slug.Deadline,
		})
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

func (r *repository) AddSegments(segments []SqlableSegmentDTO, userId string) error {
	for _, segment := range segments {
		err := r.AddSegment(segment, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) AddSegment(segment SqlableSegmentDTO, userId string) error {
	query := `INSERT INTO users_segments (user_id, segment_id, deadline) VALUES ($1, $2, $3)`

	var err error

	if segment.Deadline.String == "" {
		_, err = r.db.Exec(query, userId, segment.Id, nil)
	} else {
		_, err = r.db.Exec(query, userId, segment.Id, segment.Deadline)
	}

	if err != nil {
		return fmt.Errorf("insert to users_segments failure")
	}
	return nil
}

func (r *repository) DeleteSegments(segments []SqlableSegmentDTO, userId string) error {
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
