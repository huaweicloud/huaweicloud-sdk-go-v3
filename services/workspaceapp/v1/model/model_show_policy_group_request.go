package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyGroupRequest Request Object
type ShowPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`
}

func (o ShowPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyGroupRequest", string(data)}, " ")
}
