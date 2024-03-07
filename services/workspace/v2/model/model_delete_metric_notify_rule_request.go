package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMetricNotifyRuleRequest Request Object
type DeleteMetricNotifyRuleRequest struct {

	// 通知规则ID
	RuleId string `json:"rule_id"`
}

func (o DeleteMetricNotifyRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetricNotifyRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteMetricNotifyRuleRequest", string(data)}, " ")
}
