package model

import (
	"encoding/json"

	"strings"
)

// 实例分布对象
type Distribution struct {
	// 所在大区名称。

	Area *string `json:"area,omitempty"`
	// 所在城市名称。

	City *string `json:"city,omitempty"`
	// 所属运营商名称。

	Operator *string `json:"operator,omitempty"`
	// 所属省份名称。

	Province *string `json:"province,omitempty"`
	// 站点ID。

	SiteId *string `json:"site_id,omitempty"`
	// 线路ID。多线路场景下，将在该线路下创建弹性公网IP。

	PoolId *string `json:"pool_id,omitempty"`
	// 资源组配置模板数目

	StackCount *int32 `json:"stack_count,omitempty"`
	// 城市简称。

	CityShortName *string `json:"city_short_name,omitempty"`
}

func (o Distribution) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Distribution struct{}"
	}

	return strings.Join([]string{"Distribution", string(data)}, " ")
}
