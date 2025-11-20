package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpAssignHighRiskSysPolicyOrRoleToUserDetails 为IAM用户授予高风险系统策略或角色分析详细结果。
type IamBpAssignHighRiskSysPolicyOrRoleToUserDetails struct {

	// 用户的唯一标识符（ID）。
	UserId string `json:"user_id"`

	// 权限名称。
	PermissionName string `json:"permission_name"`
}

func (o IamBpAssignHighRiskSysPolicyOrRoleToUserDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpAssignHighRiskSysPolicyOrRoleToUserDetails struct{}"
	}

	return strings.Join([]string{"IamBpAssignHighRiskSysPolicyOrRoleToUserDetails", string(data)}, " ")
}
