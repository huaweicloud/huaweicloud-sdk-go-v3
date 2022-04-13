package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSourceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 规则ID

	RuleId string `json:"rule_id"`

	Body *CreateSourceRequestBody `json:"body,omitempty"`
}

func (o CreateSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSourceRequest struct{}"
	}

	return strings.Join([]string{"CreateSourceRequest", string(data)}, " ")
}
