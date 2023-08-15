package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceAddImageMemberRequest Request Object
type GlanceAddImageMemberRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	Body *GlanceAddImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceAddImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberRequest", string(data)}, " ")
}
