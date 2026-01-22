package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedIpsRuleVo struct {

	// 动作：0表示仅记录日志、1表示拦截session、2表示拦截ip
	Action *int32 `json:"action,omitempty"`

	// 频率ips规则id
	IpsRuleId *string `json:"ips_rule_id,omitempty"`

	// ips规则类型：0表示敏感目录扫描、1表示反弹shell
	IpsRuleType *int32 `json:"ips_rule_type,omitempty"`

	// 频率ips定义JSON字符串
	Param *string `json:"param,omitempty"`

	// 频率ips规则状态，0表示关闭，1表示开启
	Status *int32 `json:"status,omitempty"`
}

func (o AdvancedIpsRuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedIpsRuleVo struct{}"
	}

	return strings.Join([]string{"AdvancedIpsRuleVo", string(data)}, " ")
}
