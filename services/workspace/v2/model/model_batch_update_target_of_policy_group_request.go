package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTargetOfPolicyGroupRequest Request Object
type BatchUpdateTargetOfPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	Body *ModifyPolicyTargetReq `json:"body,omitempty"`
}

func (o BatchUpdateTargetOfPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTargetOfPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateTargetOfPolicyGroupRequest", string(data)}, " ")
}
