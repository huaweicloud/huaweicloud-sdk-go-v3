package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCommandMobilizationRequest Request Object
type ExecuteCommandMobilizationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ExecuteCommandRequestBody `json:"body,omitempty"`
}

func (o ExecuteCommandMobilizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCommandMobilizationRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCommandMobilizationRequest", string(data)}, " ")
}
