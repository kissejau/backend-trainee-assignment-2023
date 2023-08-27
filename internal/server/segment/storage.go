package segment

type Repository interface {
	Create(segment Segment) error
	Get(slug string) (Segment, error)
	List() ([]Segment, error)
	Update(segment Segment) error
	Delete(id string) error
}
