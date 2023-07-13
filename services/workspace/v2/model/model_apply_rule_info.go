package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyRuleInfo 认证应用对象信息，为null时代表对所有应用对象生效。
type ApplyRuleInfo struct {

	// 认证应用对象类型。 - ACCESS_MODE：接入类型
	RuleType *string `json:"rule_type,omitempty"`

	// 认证应用对象。 - INTERNET：互联网接入，rule_type为ACCESS_MODE时可选该值 - PRIVATE：专线接入，rule_type为ACCESS_MODE时可选该值
	Rule *string `json:"rule,omitempty"`
}

func (o ApplyRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyRuleInfo struct{}"
	}

	return strings.Join([]string{"ApplyRuleInfo", string(data)}, " ")
}
