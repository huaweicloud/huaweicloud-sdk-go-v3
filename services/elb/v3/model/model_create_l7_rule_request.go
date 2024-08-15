package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateL7RuleRequest Request Object
type CreateL7RuleRequest struct {

	// 参数解释：转发策略ID。
	L7policyId string `json:"l7policy_id"`

	Body *CreateL7RuleRequestBody `json:"body,omitempty"`
}

func (o CreateL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleRequest struct{}"
	}

	return strings.Join([]string{"CreateL7RuleRequest", string(data)}, " ")
}
