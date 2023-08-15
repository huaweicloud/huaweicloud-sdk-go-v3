package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpQueryCfwAccessControllerLogsResponseDtoDataRecords struct {

	// 动作0：permit,1：deny
	Action *string `json:"action,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 命中时间
	HitTime *int32 `json:"hit_time,omitempty"`

	// 文档ID
	LogId *string `json:"log_id,omitempty"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口
	SrcPort *string `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 目的端口
	DstPort *string `json:"dst_port,omitempty"`

	// 协议类型:TCP为6,UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *string `json:"protocol,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`
}

func (o HttpQueryCfwAccessControllerLogsResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwAccessControllerLogsResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwAccessControllerLogsResponseDtoDataRecords", string(data)}, " ")
}
