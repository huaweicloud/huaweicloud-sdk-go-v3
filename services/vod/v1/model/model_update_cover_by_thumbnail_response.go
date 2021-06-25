package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCoverByThumbnailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCoverByThumbnailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCoverByThumbnailResponse struct{}"
	}

	return strings.Join([]string{"UpdateCoverByThumbnailResponse", string(data)}, " ")
}
