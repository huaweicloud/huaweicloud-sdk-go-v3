package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmPoliciesEnabledStateResponse Response Object
type BatchUpdateOneClickAlarmPoliciesEnabledStateResponse struct {

	// 成功启停的告警规则策略ID列表
	AlarmPolicyIds *[]string `json:"alarm_policy_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdateOneClickAlarmPoliciesEnabledStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmPoliciesEnabledStateResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmPoliciesEnabledStateResponse", string(data)}, " ")
}
