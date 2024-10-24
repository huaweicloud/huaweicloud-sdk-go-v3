package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedManagedPolicyDto 系统身份策略详细信息
type AttachedManagedPolicyDto struct {

	// IAM系统身份策略唯一标识
	PolicyId *string `json:"policy_id,omitempty"`

	// IAM系统身份策略名称
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o AttachedManagedPolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedManagedPolicyDto struct{}"
	}

	return strings.Join([]string{"AttachedManagedPolicyDto", string(data)}, " ")
}
