package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InviteWithPwdRequest struct {
	// 会议ID

	ConferenceID string `json:"conferenceID"`

	Body *RestInviteWithPwdReqBody `json:"body,omitempty"`
}

func (o InviteWithPwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteWithPwdRequest struct{}"
	}

	return strings.Join([]string{"InviteWithPwdRequest", string(data)}, " ")
}
