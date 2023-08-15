package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleRequest Request Object
type CreateRuleRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *RuleDetail `json:"body,omitempty"`
}

func (o CreateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleRequest", string(data)}, " ")
}
