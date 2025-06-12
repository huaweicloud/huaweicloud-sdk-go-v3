package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertSrcGeo 源IP的地理位置信息
type AlertSrcGeo struct {

	// 纬度
	Latitude float32 `json:"latitude,omitempty"`

	// 经度
	Longitude float32 `json:"longitude,omitempty"`

	// 城市编码
	CityCode *string `json:"city_code,omitempty"`

	// 国家简码，参考ISO 3166-1 alpha-2，例如：CN | US | DE | IT | SG
	CountryCode *string `json:"country_code,omitempty"`
}

func (o AlertSrcGeo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertSrcGeo struct{}"
	}

	return strings.Join([]string{"AlertSrcGeo", string(data)}, " ")
}
