package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupNameInfo struct {

	// 策略组名。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o PolicyGroupNameInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupNameInfo struct{}"
	}

	return strings.Join([]string{"PolicyGroupNameInfo", string(data)}, " ")
}
