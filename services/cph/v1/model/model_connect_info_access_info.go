package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectInfoAccessInfo 云手机接入信息
type ConnectInfoAccessInfo struct {

	// 云手机实例的访问IP
	AccessIp *string `json:"access_ip,omitempty"`

	// 云手机实例的内网访问IP
	IntranetIp *string `json:"intranet_ip,omitempty"`

	// 云手机服务器IPv6 IP
	AccessIpv6 *string `json:"access_ipv6,omitempty"`

	// 云手机实例的访问端口
	AccessPort *int32 `json:"access_port,omitempty"`

	// 本次接入的会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 时间戳
	Timestamp *string `json:"timestamp,omitempty"`

	// 签名令牌
	Ticket *string `json:"ticket,omitempty"`
}

func (o ConnectInfoAccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectInfoAccessInfo struct{}"
	}

	return strings.Join([]string{"ConnectInfoAccessInfo", string(data)}, " ")
}
