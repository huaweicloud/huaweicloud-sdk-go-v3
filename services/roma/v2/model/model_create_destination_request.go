package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDestinationRequest Request Object
type CreateDestinationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	Body *CreateDestinationRequestBody `json:"body,omitempty"`
}

func (o CreateDestinationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDestinationRequest struct{}"
	}

	return strings.Join([]string{"CreateDestinationRequest", string(data)}, " ")
}
