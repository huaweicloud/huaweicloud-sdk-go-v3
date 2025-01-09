package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetOfPolicyGroupRequest Request Object
type ListTargetOfPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	// 应用对象的类型。 - INSTANCE：表示桌面。 - USER：表示用户。 - OU：表示组织单元。 - CLIENTIP：终端IP地址。
	TargetType *string `json:"target_type,omitempty"`

	// 对象名称，支持模糊查询。
	TargetName *string `json:"target_name,omitempty"`

	// 每页数量。范围0-1000
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *string `json:"offset,omitempty"`
}

func (o ListTargetOfPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetOfPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTargetOfPolicyGroupRequest", string(data)}, " ")
}
