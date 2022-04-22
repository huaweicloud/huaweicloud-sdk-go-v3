package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端状态信息
type TokenInfo struct {

	// 会控鉴权Token。
	Token *string `json:"token,omitempty"`

	// websocket建链鉴权Token，成功时必带。
	TmpWsToken *string `json:"tmpWsToken,omitempty"`

	// websocket建链URL。
	WsURL *string `json:"wsURL,omitempty"`

	// 会议中的角色 1：会议主席 0：普通与会者
	Role *int32 `json:"role,omitempty"`

	// 会话过期时间。UTC时间毫秒数。
	ExpireTime *int64 `json:"expireTime,omitempty"`

	// 会议预定人ID。
	UserID *string `json:"userID,omitempty"`

	// 会议所属企业ID。
	OrgID *string `json:"orgID,omitempty"`

	// 终端请求时，返回终端入会后会场ID。
	ParticipantID *string `json:"participantID,omitempty"`

	// 会控token失效的时间。（单位秒）
	ConfTokenExpireTime *int32 `json:"confTokenExpireTime,omitempty"`

	// 云会议室会议的当前会议ID。
	VmrCurrentConfID *string `json:"vmrCurrentConfID,omitempty"`

	// websocket消息推送支持类型。
	SupportNotifyType *[]string `json:"supportNotifyType,omitempty"`
}

func (o TokenInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenInfo struct{}"
	}

	return strings.Join([]string{"TokenInfo", string(data)}, " ")
}
