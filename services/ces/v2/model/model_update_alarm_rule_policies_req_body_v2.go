package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlarmRulePoliciesReqBodyV2 struct {

	// **参数解释**： 策略信息。 **约束限制**： 包含的策略信息数量最多为50个，最少为1个。
	Policies []UpdatePolicyReq `json:"policies"`
}

func (o UpdateAlarmRulePoliciesReqBodyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesReqBodyV2 struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesReqBodyV2", string(data)}, " ")
}
