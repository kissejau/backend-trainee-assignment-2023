package user

type SetUserSegmentsDTO struct {
	UpSlugs   []string `json:"up_slugs"`
	DownSlugs []string `json:"down_slugs"`
	UserId    string   `json:"user_id"`
}
