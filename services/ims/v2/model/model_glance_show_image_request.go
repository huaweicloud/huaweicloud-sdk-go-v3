package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceShowImageRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`
}

func (o GlanceShowImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceShowImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageRequest", string(data)}, " ")
}
