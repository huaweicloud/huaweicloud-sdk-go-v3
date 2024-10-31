package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpCcRuleRequest Request Object
type CreateHttpCcRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpCcRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpCcRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpCcRuleRequest", string(data)}, " ")
}
