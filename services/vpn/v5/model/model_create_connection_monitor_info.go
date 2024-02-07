package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectionMonitorInfo struct {

	// VPN连接监控ID
	Id *string `json:"id,omitempty"`

	// VPN连接监控对应的VPN连接ID
	VpnConnectionId *string `json:"vpn_connection_id,omitempty"`

	// 监控类型，取值范围：gateway
	Type *string `json:"type,omitempty"`

	// VPN连接监控的源地址
	SourceIp *string `json:"source_ip,omitempty"`

	// VPN连接监控的目的地址
	DestinationIp *string `json:"destination_ip,omitempty"`

	// 预留字段，nqa使用的协议类型，目前使用默认值ICMP
	ProtoType *string `json:"proto_type,omitempty"`
}

func (o CreateConnectionMonitorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionMonitorInfo struct{}"
	}

	return strings.Join([]string{"CreateConnectionMonitorInfo", string(data)}, " ")
}
