package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceDeleteImageRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`

	Body *GlanceDeleteImageRequestBody `json:"body,omitempty"`
}

func (o GlanceDeleteImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageRequest", string(data)}, " ")
}
