package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyGroupRequest Request Object
type DeletePolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`
}

func (o DeletePolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyGroupRequest", string(data)}, " ")
}
