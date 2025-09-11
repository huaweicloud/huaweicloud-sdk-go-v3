package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRulePoliciesResponse Response Object
type ListAlarmRulePoliciesResponse struct {

	// **参数解释**： 告警策略信息列表。
	Policies *[]ListPolicyResp `json:"policies,omitempty"`

	// **参数解释**： 指定告警规则对应的策略总数。 **取值范围**： 0-100
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulePoliciesResponse", string(data)}, " ")
}
