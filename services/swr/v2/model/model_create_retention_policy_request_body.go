package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRetentionPolicyRequestBody struct {

	// 算法，目前只支持or
	Algorithm string `json:"algorithm"`

	// 是否启用或者关闭所有retentionRules
	Enabled bool `json:"enabled"`

	// 匹配规则，配置repo范围、tag范围以及作用规则
	Rules []RetentionRule `json:"rules"`

	Trigger *TriggerConfig `json:"trigger"`

	// 策略名称，由字母、汉字、数字、下划线（_）、中划线 (-)组成，1-256个字符。
	Name string `json:"name"`
}

func (o CreateRetentionPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetentionPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRetentionPolicyRequestBody", string(data)}, " ")
}
