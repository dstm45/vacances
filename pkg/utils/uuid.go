package utils

import (
	"net/http"

	"github.com/google/uuid"
)

func ParseUUID(r *http.Request) (*uuid.UUID, error) {
	UUID, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		return nil, err
	}
	return &UUID, nil
}
