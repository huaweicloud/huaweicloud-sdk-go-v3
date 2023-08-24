package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionConfiguration struct {

	// TCP会话过期时间。
	TcpSessionExpireTime *int32 `json:"tcp_session_expire_time,omitempty"`

	// UDP会话过期时间。
	UdpSessionExpireTime *int32 `json:"udp_session_expire_time,omitempty"`

	// ICMP会话过期时间。
	IcmpSessionExpireTime *int32 `json:"icmp_session_expire_time,omitempty"`

	// TCP连接关闭时TIME_WAIT状态持续时间。
	TcpTimeWaitTime *int32 `json:"tcp_time_wait_time,omitempty"`
}

func (o SessionConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionConfiguration struct{}"
	}

	return strings.Join([]string{"SessionConfiguration", string(data)}, " ")
}
