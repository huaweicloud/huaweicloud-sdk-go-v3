package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceEndpointPolicyRequest Request Object
type CreateInstanceEndpointPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateEndpointPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceEndpointPolicyRequest", string(data)}, " ")
}
