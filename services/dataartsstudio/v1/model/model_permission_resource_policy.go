package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionResourcePolicy struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 资源对象列表
	Resources []ResourcePolicyItem `json:"resources"`

	// 成员列表
	Members []MemberPolicyItem `json:"members"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o PermissionResourcePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionResourcePolicy struct{}"
	}

	return strings.Join([]string{"PermissionResourcePolicy", string(data)}, " ")
}
