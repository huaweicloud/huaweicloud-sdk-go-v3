package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceUpdateImageMemberRequest struct {
	// 镜像id

	ImageId string `json:"image_id"`
	// 成员id

	MemberId string `json:"member_id"`

	Body *GlanceUpdateImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceUpdateImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberRequest", string(data)}, " ")
}
