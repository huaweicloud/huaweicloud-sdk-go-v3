package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchEnableAlarmPoliciesRequestBody struct {

	// 需要批量启停的告警规则策略的ID列表
	AlarmPolicyIds []string `json:"alarm_policy_ids"`

	// 是否启用告警策略。true:开启，false：关闭。
	Enabled bool `json:"enabled"`

	// 告警中的策略全部被停用时是否保留策略信息。true:保留；false:删除。
	RetainWhenAllDisabled *bool `json:"retain_when_all_disabled,omitempty"`
}

func (o BatchEnableAlarmPoliciesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmPoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmPoliciesRequestBody", string(data)}, " ")
}
