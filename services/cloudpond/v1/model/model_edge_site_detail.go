package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeSiteDetail 边缘小站详情
type EdgeSiteDetail struct {

	// 边缘小站ID
	Id *string `json:"id,omitempty"`

	// 边缘小站所属账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 边缘小站名称
	Name *string `json:"name,omitempty"`

	// 边缘小站描述
	Description *string `json:"description,omitempty"`

	// 边缘小站所属区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 边缘小站所属项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 边缘小站的可用区ID
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 边缘小站的部署状态
	Status *string `json:"status,omitempty"`

	Location *LocationDetail `json:"location,omitempty"`

	// [边缘小站](tag:hws)[分布式微型专属小站](tag:cmcc)所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 边缘小站创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 边缘小站更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o EdgeSiteDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeSiteDetail struct{}"
	}

	return strings.Join([]string{"EdgeSiteDetail", string(data)}, " ")
}
