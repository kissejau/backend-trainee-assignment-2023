package user

import "github.com/kissejau/backend-trainee-assignment-2023/internal/server/segment"

type Repository interface {
	Create(name string) error
	Get(id string) (User, error)
	List() ([]User, error)
	Update(u User) error
	Delete(id string) error

	GetSegments(id string) ([]segment.Segment, error)
	SetSegments(dto SetUserSegmentsDTO) error
}
