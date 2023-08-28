package user

import (
	"database/sql"
	"log"
	"time"
)

type SetUserSegmentsDTO struct {
	UpSlugs   []SegmentDTO `json:"up_slugs,omitempty"`
	DownSlugs []SegmentDTO `json:"down_slugs,omitempty"`
	UserId    string       `json:"user_id"`
}

func (set SetUserSegmentsDTO) SqlableSetUserSegmentsDTO() SqlableSetUserSegmentsDTO {
	var upSlugs, downSlugs []SqlableSegmentDTO
	sqlableSetUserSegmentsDTO := SqlableSetUserSegmentsDTO{
		UserId: set.UserId,
	}

	for _, segment := range set.UpSlugs {
		slug := SqlableSegmentDTO{
			Id:   segment.Id,
			Slug: segment.Slug,
			Deadline: sql.NullString{
				String: segment.Deadline,
			},
		}
		slug.Deadline.Scan(slug.Deadline.String)
		upSlugs = append(upSlugs, slug)
	}
	sqlableSetUserSegmentsDTO.UpSlugs = upSlugs

	for _, segment := range set.DownSlugs {
		downSlugs = append(downSlugs, SqlableSegmentDTO{
			Id:   segment.Id,
			Slug: segment.Slug,
			Deadline: sql.NullString{
				String: segment.Deadline,
			},
		})
	}
	sqlableSetUserSegmentsDTO.DownSlugs = downSlugs
	return sqlableSetUserSegmentsDTO
}

type SegmentDTO struct {
	Id       string `json:"id,omitempty"`
	Slug     string `json:"slug"`
	Deadline string `json:"deadline,omitempty"`
}

type SqlableSetUserSegmentsDTO struct {
	UpSlugs   []SqlableSegmentDTO `json:"up_slugs,omitempty"`
	DownSlugs []SqlableSegmentDTO `json:"down_slugs,omitempty"`
	UserId    string              `json:"user_id"`
}

type SqlableSegmentDTO struct {
	Id       string         `json:"id,omitempty"`
	Slug     string         `json:"slug"`
	Deadline sql.NullString `json:"deadline,omitempty"`
}

func (segmentDTO *SqlableSegmentDTO) SegmentWithActiveStatusDTO() *SegmentWithActiveStatusDTO {
	segment := &SegmentWithActiveStatusDTO{
		Id:       segmentDTO.Id,
		Slug:     segmentDTO.Slug,
		IsActive: true,
	}

	if !segmentDTO.Deadline.Valid {
		log.Println("NOTE 1")
		log.Println(segmentDTO.Deadline)
		return segment
	}
	date, _ := time.Parse("yyyy-mm-dd hh:mm:ss", segmentDTO.Deadline.String)
	if date.After(time.Now()) {
		return segment
	}
	segment.IsActive = false
	return segment
}

type SegmentWithActiveStatusDTO struct {
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	IsActive bool   `json:"is_active"`
}
