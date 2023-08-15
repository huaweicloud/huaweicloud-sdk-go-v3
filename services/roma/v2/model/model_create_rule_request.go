package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleRequest Request Object
type CreateRuleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateRuleRequestBody `json:"body,omitempty"`
}

func (o CreateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleRequest", string(data)}, " ")
}
