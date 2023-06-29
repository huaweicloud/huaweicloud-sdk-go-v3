package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceDeleteImageMemberRequest Request Object
type GlanceDeleteImageMemberRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 成员id
	MemberId string `json:"member_id"`
}

func (o GlanceDeleteImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberRequest", string(data)}, " ")
}
