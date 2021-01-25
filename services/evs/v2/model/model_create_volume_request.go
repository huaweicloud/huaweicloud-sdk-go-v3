package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVolumeRequest struct {
	Body *CreateVolumeRequestBody `json:"body,omitempty"`
}

func (o CreateVolumeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVolumeRequest struct{}"
	}

	return strings.Join([]string{"CreateVolumeRequest", string(data)}, " ")
}
