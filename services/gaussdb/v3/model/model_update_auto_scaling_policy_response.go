package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScalingPolicyResponse Response Object
type UpdateAutoScalingPolicyResponse struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	SwitchStatus   *AutoScalingSwitchStatus `json:"switch_status,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoScalingPolicyResponse", string(data)}, " ")
}
