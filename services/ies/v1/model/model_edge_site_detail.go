package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘小站详情
type EdgeSiteDetail struct {

	// 边缘小站ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 边缘小站所属账号ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 边缘小站名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 边缘小站描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 边缘小站所属区域ID
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 边缘小站所属项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 边缘小站的可用区ID
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty" xml:"availability_zone_id"`

	// 边缘小站的部署状态
	Status *string `json:"status,omitempty" xml:"status"`

	Location *LocationDetail `json:"location,omitempty" xml:"location"`

	// 边缘小站创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty" xml:"created_at"`

	// 边缘小站更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty" xml:"updated_at"`
}

func (o EdgeSiteDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeSiteDetail struct{}"
	}

	return strings.Join([]string{"EdgeSiteDetail", string(data)}, " ")
}
