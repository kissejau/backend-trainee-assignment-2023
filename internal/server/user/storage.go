package user

type Repository interface {
	Create(name string) error
	Get(id string) (User, error)
	List() ([]User, error)
	Update(u User) error
	Delete(id string) error

	GetSegments(id string) ([]SegmentWithActiveStatusDTO, error)
	SetSegments(dto SqlableSetUserSegmentsDTO) error
}
