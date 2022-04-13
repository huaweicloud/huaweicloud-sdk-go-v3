package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 登录响应
type RestInviteWithPwdReqBody struct {
	// 被邀请的会议人号码。

	CallNum string `json:"callNum"`
	// 被邀请的会议人所属企业ID。

	OrgID *string `json:"orgID,omitempty"`
	// 当前会议ID。

	ConfID string `json:"confID"`
	// 当前会议的密码。

	Pwd string `json:"pwd"`
	// 号码类型0是本局号码，1是中继号码。

	NumBelongsType *int32 `json:"numBelongsType,omitempty"`
	// 是否不叠加会场名。

	IsNotOverlayPidName *bool `json:"isNotOverlayPidName,omitempty"`
}

func (o RestInviteWithPwdReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestInviteWithPwdReqBody struct{}"
	}

	return strings.Join([]string{"RestInviteWithPwdReqBody", string(data)}, " ")
}
