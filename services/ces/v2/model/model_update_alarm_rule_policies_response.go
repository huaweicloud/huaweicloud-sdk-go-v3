package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRulePoliciesResponse Response Object
type UpdateAlarmRulePoliciesResponse struct {

	// **参数解释**： 策略信息。
	Policies       *[]UpdatePolicyResp `json:"policies,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesResponse", string(data)}, " ")
}
