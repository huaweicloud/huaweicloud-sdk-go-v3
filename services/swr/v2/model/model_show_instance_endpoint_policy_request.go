package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceEndpointPolicyRequest Request Object
type ShowInstanceEndpointPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceEndpointPolicyRequest", string(data)}, " ")
}
