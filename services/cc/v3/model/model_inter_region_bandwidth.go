package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 域间带宽实例。
type InterRegionBandwidth struct {

	// 域间带宽实例的ID。
	Id *string `json:"id,omitempty"`

	// 域间带宽实例的名字。
	Name *string `json:"name,omitempty"`

	// 域间带宽实例的描述。
	Description *string `json:"description,omitempty"`

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 带宽包实例的ID。
	BandwidthPackageId *string `json:"bandwidth_package_id,omitempty"`

	// 域间带宽实例的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 域间带宽实例的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 云连接实例的ID。
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// 域间实例信息。
	InterRegions *[]InterRegion `json:"inter_regions,omitempty"`

	// 域间带宽的值。
	Bandwidth *int32 `json:"bandwidth,omitempty"`
}

func (o InterRegionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterRegionBandwidth struct{}"
	}

	return strings.Join([]string{"InterRegionBandwidth", string(data)}, " ")
}
