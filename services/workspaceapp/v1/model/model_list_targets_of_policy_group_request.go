package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetsOfPolicyGroupRequest Request Object
type ListTargetsOfPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	// 应用对象的类型： - USER：表示用户。 - USERGROUP：表示用户组。 - APPGROUP：应用组。 - OU：组织单元。 - ALl：所有类型
	TargetType *string `json:"target_type,omitempty"`
}

func (o ListTargetsOfPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetsOfPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTargetsOfPolicyGroupRequest", string(data)}, " ")
}
