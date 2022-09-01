package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAvailableZonesRespAvailableZones struct {

	// 是否售罄。
	SoldOut *bool `json:"soldOut,omitempty" xml:"soldOut"`

	// 可用区ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 可用区编码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 可用区名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 可用区端口号。
	Port *string `json:"port,omitempty" xml:"port"`

	// 分区上是否还有可用资源。
	ResourceAvailability *string `json:"resource_availability,omitempty" xml:"resource_availability"`

	// 是否为默认可用区。
	DefaultAz *bool `json:"default_az,omitempty" xml:"default_az"`

	// 剩余时间。
	RemainTime *int64 `json:"remain_time,omitempty" xml:"remain_time"`

	// 是否支持IPv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty" xml:"ipv6_enable"`
}

func (o ListAvailableZonesRespAvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesRespAvailableZones struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesRespAvailableZones", string(data)}, " ")
}
