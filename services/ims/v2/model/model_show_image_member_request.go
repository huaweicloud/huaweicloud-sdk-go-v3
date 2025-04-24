package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageMemberRequest Request Object
type ShowImageMemberRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 成员id
	MemberId string `json:"member_id"`
}

func (o ShowImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowImageMemberRequest", string(data)}, " ")
}
