package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type GeoLocation struct {

	// 地理位置信息ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 所在大区。
	Area *string `json:"area,omitempty" xml:"area"`

	// 所在城市。
	City *string `json:"city,omitempty" xml:"city"`

	// 所在的国家。
	Country *string `json:"country,omitempty" xml:"country"`

	// 区域的国际化名称。
	I18nArea *string `json:"i18n_area,omitempty" xml:"i18n_area"`

	// 城市的国际化名称。
	I18nCity *string `json:"i18n_city,omitempty" xml:"i18n_city"`

	// 国家的国际化名称。
	I18nCountry *string `json:"i18n_country,omitempty" xml:"i18n_country"`

	// 省份的国际化名称。
	I18nProvince *string `json:"i18n_province,omitempty" xml:"i18n_province"`

	// 所在省份。
	Province *string `json:"province,omitempty" xml:"province"`
}

func (o GeoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeoLocation struct{}"
	}

	return strings.Join([]string{"GeoLocation", string(data)}, " ")
}
