package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedIpsRuleDto struct {

	// 动作：0表示仅记录日志、1表示拦截session、2表示拦截ip
	Action *int32 `json:"action,omitempty"`

	// 高级ips规则id
	IpsRuleId *string `json:"ips_rule_id,omitempty"`

	// ips规则类型：0表示敏感目录扫描、1表示反弹xshell
	IpsRuleType *int32 `json:"ips_rule_type,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 包含特殊参数的JSON字符串
	Param *string `json:"param,omitempty"`

	// 开关状态：0表示关闭、1表示开启
	Status *int32 `json:"status,omitempty"`
}

func (o AdvancedIpsRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedIpsRuleDto struct{}"
	}

	return strings.Join([]string{"AdvancedIpsRuleDto", string(data)}, " ")
}
