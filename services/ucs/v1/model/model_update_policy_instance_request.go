package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyInstanceRequest Request Object
type UpdatePolicyInstanceRequest struct {

	// 策略实例id
	Policyinstanceid string `json:"policyinstanceid"`

	Body *UcsConstraintRequest `json:"body,omitempty"`
}

func (o UpdatePolicyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyInstanceRequest", string(data)}, " ")
}
