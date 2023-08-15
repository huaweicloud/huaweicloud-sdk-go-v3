package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestInviteReqBody 邀请与会者请求。
type RestInviteReqBody struct {

	// 邀请的与会者列表。
	Attendees []Attendee `json:"attendees"`
}

func (o RestInviteReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestInviteReqBody struct{}"
	}

	return strings.Join([]string{"RestInviteReqBody", string(data)}, " ")
}
