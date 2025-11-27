package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyInstanceRequest Request Object
type ShowPolicyInstanceRequest struct {

	// 策略实例id
	Policyinstanceid string `json:"policyinstanceid"`
}

func (o ShowPolicyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyInstanceRequest", string(data)}, " ")
}
