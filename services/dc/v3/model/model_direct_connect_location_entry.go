package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DirectConnectLocationEntry direct-connect-location专线接入点详情
type DirectConnectLocationEntry struct {

	// 专线接入站点资源的ID
	Id *string `json:"id,omitempty"`

	// 专线接入点的名称
	Name *string `json:"name,omitempty"`

	// Location所属Region
	RegionId *string `json:"region_id,omitempty"`

	// 专线接入点对应的站点编码
	SiteCode *string `json:"site_code,omitempty"`

	Address *AddressBody `json:"address,omitempty"`

	Locales *LocalesBody `json:"locales,omitempty"`

	// 可支持的运营商列表。
	ProviderList *[]ProviderResponseBody `json:"provider_list,omitempty"`

	// 专线接入点所属public_border_group
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 地理位置纬度
	Latitude *string `json:"latitude,omitempty"`

	// 地理位置经度
	Longitude *string `json:"longitude,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 创建时间。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间。
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 接入点内设备可选的端口类型
	AvailablePortSpeeds *[]string `json:"available_port_speeds,omitempty"`
}

func (o DirectConnectLocationEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectConnectLocationEntry struct{}"
	}

	return strings.Join([]string{"DirectConnectLocationEntry", string(data)}, " ")
}
