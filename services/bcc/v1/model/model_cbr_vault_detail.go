package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbrVaultDetail struct {

	// 存储库ID
	Id *string `json:"id,omitempty"`

	// 存储库名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 存储库提供云服务，CBR
	Provider *string `json:"provider,omitempty"`

	// 地域id
	RegionId *string `json:"region_id,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目id，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 创建时间
	Created *string `json:"created,omitempty"`

	ResourceDetail *CbrVaultDetailResourceDetail `json:"resource_detail,omitempty"`

	// 资源标签
	ResourceLabel *[]Tag `json:"resource_label,omitempty"`
}

func (o CbrVaultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbrVaultDetail struct{}"
	}

	return strings.Join([]string{"CbrVaultDetail", string(data)}, " ")
}
