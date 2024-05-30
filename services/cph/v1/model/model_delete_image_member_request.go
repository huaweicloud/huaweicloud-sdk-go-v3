package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageMemberRequest Request Object
type DeleteImageMemberRequest struct {

	// 镜像id。
	ImageId string `json:"image_id"`

	// 被共享账号的PROJECT_ID。
	MemberId string `json:"member_id"`
}

func (o DeleteImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageMemberRequest", string(data)}, " ")
}
