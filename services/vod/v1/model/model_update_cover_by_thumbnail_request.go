package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateCoverByThumbnailRequest struct {
	Body *UpdateCoverByThumbnailReq `json:"body,omitempty"`
}

func (o UpdateCoverByThumbnailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCoverByThumbnailRequest struct{}"
	}

	return strings.Join([]string{"UpdateCoverByThumbnailRequest", string(data)}, " ")
}
