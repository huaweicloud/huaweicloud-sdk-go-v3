package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleServiceDto RuleServiceDto
type RuleServiceDto struct {

	// 服务输入类型，0为手动输入类型，1为自动输入类型
	Type int32 `json:"type"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *int32 `json:"protocol,omitempty"`

	// 源端口
	SourcePort *string `json:"source_port,omitempty"`

	// 目的端口
	DestPort *string `json:"dest_port,omitempty"`

	// 服务组id，手动类型为空，自动类型为非空
	ServiceSetId *string `json:"service_set_id,omitempty"`

	// 服务组名称
	ServiceSetName *string `json:"service_set_name,omitempty"`
}

func (o RuleServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleServiceDto struct{}"
	}

	return strings.Join([]string{"RuleServiceDto", string(data)}, " ")
}
