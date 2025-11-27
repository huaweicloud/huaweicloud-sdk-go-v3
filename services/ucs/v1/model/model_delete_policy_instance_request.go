package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyInstanceRequest Request Object
type DeletePolicyInstanceRequest struct {

	// 策略实例id
	Policyinstanceid string `json:"policyinstanceid"`
}

func (o DeletePolicyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyInstanceRequest", string(data)}, " ")
}
