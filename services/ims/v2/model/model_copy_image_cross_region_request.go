package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CopyImageCrossRegionRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`

	Body *CopyImageCrossRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageCrossRegionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequest", string(data)}, " ")
}
