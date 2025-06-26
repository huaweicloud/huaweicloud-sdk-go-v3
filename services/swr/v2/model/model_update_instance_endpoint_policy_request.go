package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceEndpointPolicyRequest Request Object
type UpdateInstanceEndpointPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateWhiteListRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceEndpointPolicyRequest", string(data)}, " ")
}
