package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateRequestThrottlingPolicyV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *ThrottleApiBindingCreate `json:"body,omitempty"`
}

func (o AssociateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
