package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPoliciesRequest Request Object
type ListAiPoliciesRequest struct {

	// **参数解释**： 策略组ID **约束限制**： 不涉及 **取值范围**： 字符长度1-20位 **默认取值**： 不涉及
	PolicyGroupId string `json:"policy_group_id"`
}

func (o ListAiPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListAiPoliciesRequest", string(data)}, " ")
}
