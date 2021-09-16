package model

import (
	"encoding/json"

	"strings"
)

// 覆盖策略对象
type CoverageResp struct {
	// 区域调度策略，只支持centralize/discrete。  - centralize：代表城市集中策略，指定该策略，边缘业务创建时会保证将所有实例都发放在同一个站点。 - discrete：代表城市分散，指定该策略，边缘业务创建时，尽量保证所有实例都分散发放在不同站点。

	CoveragePolicy string `json:"coverage_policy"`
	// 区域分布层级，只支持area/prov/city/site。  - area:大区，用户的资源会在指定大区下发放。 - prov:省份，用户的资源会在指定省份下发放。 - city:城市，用户的资源会在指定城市下发放。 - site:站点级别。  约束：  站点层级，coverage_policy上仅支持'centralize'，coverage_sites中'site'字段仅支持使用ID(站点ID，通过“查询边缘站点列表”获取)，不支持name。

	CoverageLevel string `json:"coverage_level"`
	// 区域及购买数量列表。

	CoverageSites []CoverageSiteResp `json:"coverage_sites"`
}

func (o CoverageResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CoverageResp struct{}"
	}

	return strings.Join([]string{"CoverageResp", string(data)}, " ")
}
