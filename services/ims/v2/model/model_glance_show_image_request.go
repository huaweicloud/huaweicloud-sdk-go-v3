package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceShowImageRequest struct {

	// 镜像ID
	ImageId string `json:"image_id" xml:"image_id"`
}

func (o GlanceShowImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageRequest", string(data)}, " ")
}
