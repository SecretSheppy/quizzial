package models

import (
	"errors"
	"github.com/google/uuid"
)

type SectionResource struct {
	Resource
	SectionID uuid.UUID `gorm:"foreignKey:SectionID"`
}

func NewSectionResource(path, alt string, sID uuid.UUID) (*SectionResource, error) {
	if path == "" {
		return nil, errors.New("path must not be empty")
	}

	if alt == "" {
		return nil, errors.New("alt must not be empty")
	}

	if sID == uuid.Nil {
		return nil, errors.New("sID must not be empty")
	}

	return &SectionResource{
		Resource: Resource{
			ResourceID: uuid.New(),
			Path:       path,
			Alt:        alt,
		},
		SectionID: sID,
	}, nil
}
