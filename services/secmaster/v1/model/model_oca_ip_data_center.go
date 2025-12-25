package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpDataCenter 线下机房数据中心
type OcaIpDataCenter struct {

	// 纬度
	Latitude *float32 `json:"latitude,omitempty"`

	// 经度
	Longitude *float32 `json:"longitude,omitempty"`

	// 城市编码
	CityCode string `json:"city_code"`

	// 国家编码
	CountryCode string `json:"country_code"`
}

func (o OcaIpDataCenter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpDataCenter struct{}"
	}

	return strings.Join([]string{"OcaIpDataCenter", string(data)}, " ")
}
