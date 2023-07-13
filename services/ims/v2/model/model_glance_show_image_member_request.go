package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceShowImageMemberRequest Request Object
type GlanceShowImageMemberRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 成员id
	MemberId string `json:"member_id"`
}

func (o GlanceShowImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberRequest", string(data)}, " ")
}
