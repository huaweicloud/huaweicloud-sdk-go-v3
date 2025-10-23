package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbrPolicyDetail CbrPolicyDetail
type CbrPolicyDetail struct {

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源归属云服务
	Provider *string `json:"provider,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签
	ResourceLabel *[]Tag `json:"resource_label,omitempty"`

	ResourceDetail *CbrPolicyDetailResourceDetail `json:"resource_detail,omitempty"`
}

func (o CbrPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbrPolicyDetail struct{}"
	}

	return strings.Join([]string{"CbrPolicyDetail", string(data)}, " ")
}
