package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDisassociateThrottlingPolicyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 必须为delete
	Action string `json:"action" xml:"action"`

	Body *ThrottleBindingBatchDelete `json:"body,omitempty" xml:"body"`
}

func (o BatchDisassociateThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateThrottlingPolicyV2Request", string(data)}, " ")
}
