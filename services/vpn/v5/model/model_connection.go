package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Connection struct {

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 客户端虚拟IP
	ClientVirtualIp *string `json:"client_virtual_ip,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端用户名称
	ClientUserName *string `json:"client_user_name,omitempty"`

	// 入网包数
	InboundPackets *int64 `json:"inbound_packets,omitempty"`

	// 出网包数
	OutboundPackets *int64 `json:"outbound_packets,omitempty"`

	// 入网字节数
	InboundBytes *int64 `json:"inbound_bytes,omitempty"`

	// 出网字节数
	OutboundBytes *int64 `json:"outbound_bytes,omitempty"`

	// 连接建立时间
	ConnectionEstablishedTime *sdktime.SdkTime `json:"connection_established_time,omitempty"`

	// 时间戳
	Timestamp *sdktime.SdkTime `json:"timestamp,omitempty"`
}

func (o Connection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connection struct{}"
	}

	return strings.Join([]string{"Connection", string(data)}, " ")
}
