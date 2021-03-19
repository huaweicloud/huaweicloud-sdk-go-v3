package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CinderExportToImageRequest struct {
	VolumeId string `json:"volume_id"`

	Body *CinderExportToImageRequestBody `json:"body,omitempty"`
}

func (o CinderExportToImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderExportToImageRequest struct{}"
	}

	return strings.Join([]string{"CinderExportToImageRequest", string(data)}, " ")
}
