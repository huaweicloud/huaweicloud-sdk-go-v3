package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例分布对象
type Distribution struct {

	// 所在大区名称。
	Area *string `json:"area,omitempty" xml:"area"`

	// 所在城市名称。
	City *string `json:"city,omitempty" xml:"city"`

	// 所属运营商名称。
	Operator *string `json:"operator,omitempty" xml:"operator"`

	// 所属省份名称。
	Province *string `json:"province,omitempty" xml:"province"`

	// 站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`

	// 线路ID。多线路场景下，将在该线路下创建弹性公网IP。
	PoolId *string `json:"pool_id,omitempty" xml:"pool_id"`

	// 资源组配置模板数目
	StackCount *int32 `json:"stack_count,omitempty" xml:"stack_count"`

	// 城市简称。
	CityShortName *string `json:"city_short_name,omitempty" xml:"city_short_name"`
}

func (o Distribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Distribution struct{}"
	}

	return strings.Join([]string{"Distribution", string(data)}, " ")
}
