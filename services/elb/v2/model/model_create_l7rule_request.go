package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateL7ruleRequest struct {
	// 转发策略id

	L7policyId string `json:"l7policy_id"`

	Body *CreateL7ruleRequestBody `json:"body,omitempty"`
}

func (o CreateL7ruleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleRequest struct{}"
	}

	return strings.Join([]string{"CreateL7ruleRequest", string(data)}, " ")
}
