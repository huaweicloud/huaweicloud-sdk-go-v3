package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityResourcePermissionPolicyRequest Request Object
type ShowSecurityResourcePermissionPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 空间资源权限策略id。
	PolicyId string `json:"policy_id"`
}

func (o ShowSecurityResourcePermissionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityResourcePermissionPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityResourcePermissionPolicyRequest", string(data)}, " ")
}
