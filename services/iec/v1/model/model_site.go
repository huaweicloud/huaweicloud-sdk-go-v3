package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘站点对象类型
type Site struct {

	// 边缘站点ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 边缘站点名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 站点所在城市。
	City *string `json:"city,omitempty" xml:"city"`

	// 城市的国际化名称。
	I18nCity *string `json:"i18n_city,omitempty" xml:"i18n_city"`

	// 站点所在省份。
	Province *string `json:"province,omitempty" xml:"province"`

	// 省份的国际化名称。
	I18nProvince *string `json:"i18n_province,omitempty" xml:"i18n_province"`

	// 所在大区。
	Area *string `json:"area,omitempty" xml:"area"`

	// 大区的国际化名称。
	I18nArea *string `json:"i18n_area,omitempty" xml:"i18n_area"`

	// 站点所在的国家。
	Country *string `json:"country,omitempty" xml:"country"`

	// 国家的国际化名称。
	I18nCountry *string `json:"i18n_country,omitempty" xml:"i18n_country"`

	// 站点当前的状态。  取值范围： - Normal(正常商用) - Obt(公测) - Gray(灰度) - Offline(下线) - Promotion(推荐，也是商用) - sellout(售罄)
	Status *string `json:"status,omitempty" xml:"status"`

	// 站点IP线路列表。
	Pools *[]IpPool `json:"pools,omitempty" xml:"pools"`

	// 城市名称缩写。
	CityShortName *string `json:"city_short_name,omitempty" xml:"city_short_name"`
}

func (o Site) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Site struct{}"
	}

	return strings.Join([]string{"Site", string(data)}, " ")
}
