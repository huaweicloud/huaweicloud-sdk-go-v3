package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EwProtectResourceInfo 东西向防护的资源信息，例如VPC、VGW等
type EwProtectResourceInfo struct {

	// 防护资源类型：0 VPC，1 VGW
	ProtectedResourceType int32 `json:"protected_resource_type"`

	// 防护资源名称
	ProtectedResourceName string `json:"protected_resource_name"`

	// 防护资源id
	ProtectedResourceId string `json:"protected_resource_id"`

	// 防护资源nat网关名称
	ProtectedResourceNatName *string `json:"protected_resource_nat_name,omitempty"`

	// 防护资源nat网关id
	ProtectedResourceNatId *string `json:"protected_resource_nat_id,omitempty"`

	// 防护资源租户id
	ProtectedResourceProjectId *string `json:"protected_resource_project_id,omitempty"`
}

func (o EwProtectResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EwProtectResourceInfo struct{}"
	}

	return strings.Join([]string{"EwProtectResourceInfo", string(data)}, " ")
}
