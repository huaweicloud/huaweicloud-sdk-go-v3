package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateImageRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`

	Body *[]UpdateImageRequestBody `json:"body,omitempty"`
}

func (o UpdateImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateImageRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageRequest", string(data)}, " ")
}
