package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityResourcePermissionPolicyResponse Response Object
type ShowSecurityResourcePermissionPolicyResponse struct {

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 资源对象列表
	Resources *[]ResourcePolicyItem `json:"resources,omitempty"`

	// 成员列表
	Members *[]MemberPolicyItem `json:"members,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 修改时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSecurityResourcePermissionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityResourcePermissionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityResourcePermissionPolicyResponse", string(data)}, " ")
}
