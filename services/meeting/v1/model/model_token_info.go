package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会控Token信息。
type TokenInfo struct {

	// 会控Token。有效期半个小时。
	Token *string `json:"token,omitempty"`

	// 会控WebSocket建链鉴权Token。
	TmpWsToken *string `json:"tmpWsToken,omitempty"`

	// 会控WebSocket建链URL。
	WsURL *string `json:"wsURL,omitempty"`

	// 会议中的角色。 * 0 ：普通与会者 * 1 ：会议主持人
	Role *int32 `json:"role,omitempty"`

	// 会控Token过期时间戳（单位：毫秒）。
	ExpireTime *int64 `json:"expireTime,omitempty"`

	// 会议预定者的用户UUID。
	UserID *string `json:"userID,omitempty"`

	// 会议所属企业ID。
	OrgID *string `json:"orgID,omitempty"`

	// 终端请求时，返回终端入会后会场ID。 > 该参数将废弃，请勿使用。
	ParticipantID *string `json:"participantID,omitempty"`

	// 会控Token有效时长（单位秒）。
	ConfTokenExpireTime *int32 `json:"confTokenExpireTime,omitempty"`

	// 云会议室会议的当前会议ID。
	VmrCurrentConfID *string `json:"vmrCurrentConfID,omitempty"`

	// 会控WebSocket消息推送支持类型。
	SupportNotifyType *[]string `json:"supportNotifyType,omitempty"`
}

func (o TokenInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenInfo struct{}"
	}

	return strings.Join([]string{"TokenInfo", string(data)}, " ")
}
