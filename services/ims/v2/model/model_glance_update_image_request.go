package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceUpdateImageRequest struct {

	// 镜像ID
	ImageId string `json:"image_id" xml:"image_id"`

	Body *[]GlanceUpdateImageRequestBody `json:"body,omitempty" xml:"body"`
}

func (o GlanceUpdateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageRequest", string(data)}, " ")
}
