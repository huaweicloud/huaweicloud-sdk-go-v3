package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionMonitorRequestBodyContent 创建连接监控请求体
type CreateConnectionMonitorRequestBodyContent struct {

	// VPN连接监控对应的VPN连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o CreateConnectionMonitorRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionMonitorRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateConnectionMonitorRequestBodyContent", string(data)}, " ")
}
