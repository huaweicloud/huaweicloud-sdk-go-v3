package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails 为IAM委托授予高风险系统权限或角色分析详细结果。
type IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails struct {

	// 委托的唯一标识符（ID）。
	AgencyId string `json:"agency_id"`

	// 权限名称。
	PermissionName string `json:"permission_name"`
}

func (o IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails struct{}"
	}

	return strings.Join([]string{"IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails", string(data)}, " ")
}
