package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableZonesResponse Response Object
type ListAvailableZonesResponse struct {

	// 可用区列表
	AvailableZones *[]AvailableZoneV2 `json:"available_zones,omitempty"`

	// 默认可用区编码
	DefaultAzCode *string `json:"default_az_code,omitempty"`

	// 支持的物理可用区分组
	SupportPhysicalAzGroup *bool `json:"support_physical_az_group,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
