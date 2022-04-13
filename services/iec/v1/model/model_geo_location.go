package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type GeoLocation struct {
	// 地理位置信息ID。

	Id *string `json:"id,omitempty"`
	// 所在大区。

	Area *string `json:"area,omitempty"`
	// 所在城市。

	City *string `json:"city,omitempty"`
	// 所在的国家。

	Country *string `json:"country,omitempty"`
	// 区域的国际化名称。

	I18nArea *string `json:"i18n_area,omitempty"`
	// 城市的国际化名称。

	I18nCity *string `json:"i18n_city,omitempty"`
	// 国家的国际化名称。

	I18nCountry *string `json:"i18n_country,omitempty"`
	// 省份的国际化名称。

	I18nProvince *string `json:"i18n_province,omitempty"`
	// 所在省份。

	Province *string `json:"province,omitempty"`
}

func (o GeoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeoLocation struct{}"
	}

	return strings.Join([]string{"GeoLocation", string(data)}, " ")
}
