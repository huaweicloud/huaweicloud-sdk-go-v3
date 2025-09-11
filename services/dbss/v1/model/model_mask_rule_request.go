package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MaskRuleRequest struct {

	// 匹配正则
	MaskRegex string `json:"mask_regex"`

	// 脱敏值
	MaskValue string `json:"mask_value"`

	// 隐私数据保护规则名称
	RuleName string `json:"rule_name"`
}

func (o MaskRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaskRuleRequest struct{}"
	}

	return strings.Join([]string{"MaskRuleRequest", string(data)}, " ")
}
