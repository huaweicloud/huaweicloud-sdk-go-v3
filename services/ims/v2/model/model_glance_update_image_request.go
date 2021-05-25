package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceUpdateImageRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`

	Body *[]GlanceUpdateImageRequestBody `json:"body,omitempty"`
}

func (o GlanceUpdateImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageRequest", string(data)}, " ")
}
