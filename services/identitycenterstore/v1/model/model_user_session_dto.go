package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserSessionDto 用户session详情
type UserSessionDto struct {

	// 会话创建时间
	CreationTime int64 `json:"creation_time"`

	// 用户IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 会话ID
	SessionId string `json:"session_id"`

	// 会话失效时间
	SessionNotValidAfter int64 `json:"session_not_valid_after"`

	// 用户客户端信息
	UserAgent *string `json:"user_agent,omitempty"`

	// 用户唯一标识ID
	UserId string `json:"user_id"`
}

func (o UserSessionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSessionDto struct{}"
	}

	return strings.Join([]string{"UserSessionDto", string(data)}, " ")
}
