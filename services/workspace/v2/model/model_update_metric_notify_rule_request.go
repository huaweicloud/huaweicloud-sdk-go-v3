package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetricNotifyRuleRequest Request Object
type UpdateMetricNotifyRuleRequest struct {

	// 通知规则ID
	RuleId string `json:"rule_id"`

	Body *UpdateMetricNotifyRuleReq `json:"body,omitempty"`
}

func (o UpdateMetricNotifyRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetricNotifyRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateMetricNotifyRuleRequest", string(data)}, " ")
}
