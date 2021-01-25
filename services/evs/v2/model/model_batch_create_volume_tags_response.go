package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateVolumeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateVolumeTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsResponse", string(data)}, " ")
}
