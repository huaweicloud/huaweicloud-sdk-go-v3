package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘站点对象类型
type Site struct {
	// 边缘站点ID。

	Id *string `json:"id,omitempty"`
	// 边缘站点名称。

	Name *string `json:"name,omitempty"`
	// 站点所在城市。

	City *string `json:"city,omitempty"`
	// 城市的国际化名称。

	I18nCity *string `json:"i18n_city,omitempty"`
	// 站点所在省份。

	Province *string `json:"province,omitempty"`
	// 省份的国际化名称。

	I18nProvince *string `json:"i18n_province,omitempty"`
	// 所在大区。

	Area *string `json:"area,omitempty"`
	// 大区的国际化名称。

	I18nArea *string `json:"i18n_area,omitempty"`
	// 站点所在的国家。

	Country *string `json:"country,omitempty"`
	// 国家的国际化名称。

	I18nCountry *string `json:"i18n_country,omitempty"`
	// 站点当前的状态。  取值范围： - Normal(正常商用) - Obt(公测) - Gray(灰度) - Offline(下线) - Promotion(推荐，也是商用) - sellout(售罄)

	Status *string `json:"status,omitempty"`
	// 站点IP线路列表。

	Pools *[]IpPool `json:"pools,omitempty"`
	// 城市名称缩写。

	CityShortName *string `json:"city_short_name,omitempty"`
}

func (o Site) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Site struct{}"
	}

	return strings.Join([]string{"Site", string(data)}, " ")
}
