package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransferRuleInfo struct {

	// 转发规则名ID
	RuleId *string `json:"rule_id,omitempty"`

	// 转发协议，tcp或udp
	ForwardProtocol *string `json:"forward_protocol,omitempty"`

	// 转发端口
	ForwardPort *int32 `json:"forward_port,omitempty"`

	// 源站端口
	SourcePort *int32 `json:"source_port,omitempty"`

	// LVS转发规则
	LbMethod *string `json:"lb_method,omitempty"`

	// 源站IP，多个IP用逗号隔开，限制20个IP
	SourceIp *string `json:"source_ip,omitempty"`

	// 源站状态 1 正常，2 异常
	Status *int32 `json:"status,omitempty"`
}

func (o TransferRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferRuleInfo struct{}"
	}

	return strings.Join([]string{"TransferRuleInfo", string(data)}, " ")
}
