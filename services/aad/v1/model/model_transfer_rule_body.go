package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransferRuleBody struct {

	// 转发协议，tcp或udp
	ForwardProtocol *string `json:"forward_protocol,omitempty"`

	// 转发端口
	ForwardPort *int32 `json:"forward_port,omitempty"`

	// 源站端口
	SourcePort *int32 `json:"source_port,omitempty"`

	// 源站IP，多个IP用逗号隔开，限制20个IP
	SourceIp *string `json:"source_ip,omitempty"`
}

func (o TransferRuleBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferRuleBody struct{}"
	}

	return strings.Join([]string{"TransferRuleBody", string(data)}, " ")
}
