package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestInviteWithPwdReqBody 会议信息。
type RestInviteWithPwdReqBody struct {

	// 被邀请者的SIP号码。
	CallNum string `json:"callNum"`

	// 被邀请者所属企业ID。
	OrgID *string `json:"orgID,omitempty"`

	// 会议ID。
	ConfID string `json:"confID"`

	// 会议的密码（主持人或者来宾）。
	Pwd string `json:"pwd"`

	// SIP号码类型。默认是0。 * 0：华为云会议的号码 * 1：VC会议的号码
	NumBelongsType *int32 `json:"numBelongsType,omitempty"`

	// 是否不叠加会场名（VDC场景下适用）。
	IsNotOverlayPidName *bool `json:"isNotOverlayPidName,omitempty"`
}

func (o RestInviteWithPwdReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestInviteWithPwdReqBody struct{}"
	}

	return strings.Join([]string{"RestInviteWithPwdReqBody", string(data)}, " ")
}
