package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WksEdgeSiteDetail 边缘小站信息。
type WksEdgeSiteDetail struct {

	// 边缘小站id。
	Id *string `json:"id,omitempty"`

	// 边缘小站名称。
	Name *string `json:"name,omitempty"`

	// 边缘小站描述。
	Description *string `json:"description,omitempty"`

	// 边缘小站的可用区ID。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 部署位置。
	Address *string `json:"address,omitempty"`

	// 边缘小站的部署状态。 - initial：待部署。 - deploying：部署中。 - available：可用。 - unavailable：不可用。
	Status *string `json:"status,omitempty"`

	// 边缘小站创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 边缘小站更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// cpu使用率。
	CpuUsageRate *float64 `json:"cpu_usage_rate,omitempty"`

	// 内存使用率。
	MemoryUsageRate *float64 `json:"memory_usage_rate,omitempty"`

	// 存储使用率。
	CapacityUsageRate *float64 `json:"capacity_usage_rate,omitempty"`

	SiteInfo *SimpleSiteInfo `json:"site_info,omitempty"`
}

func (o WksEdgeSiteDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WksEdgeSiteDetail struct{}"
	}

	return strings.Join([]string{"WksEdgeSiteDetail", string(data)}, " ")
}
