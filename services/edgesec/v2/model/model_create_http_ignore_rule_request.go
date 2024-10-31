package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpIgnoreRuleRequest Request Object
type CreateHttpIgnoreRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpIgnoreRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpIgnoreRuleRequest", string(data)}, " ")
}
