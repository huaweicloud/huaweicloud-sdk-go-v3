package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRequest struct {

	// 批量启停的告警规则列表。批量启停告警规则时，该参数必填。
	AlarmRules *[]BatchAlarmRulesBody `json:"alarm_rules,omitempty"`

	// 批量修改告警行动规则的告警规则列表。批量修改告警行动规则时，该参数必填。
	UpdateActionRules *[]BatchUpdateActionRules `json:"update_action_rules,omitempty"`

	// 更新类型：BATCH_UPDATE_ACTION_RULE。批量修改告警行动规则时，该参数必填。
	UpdateType *string `json:"update_type,omitempty"`
}

func (o BatchUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateRequest", string(data)}, " ")
}
