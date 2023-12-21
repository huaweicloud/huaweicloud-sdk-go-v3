package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CoverageInstance 覆盖策略对象
type CoverageInstance struct {

	// 区域调度策略，只支持centralize。  - centralize：代表城市集中策略，指定该策略，边缘业务创建时会保证将所有实例都发放在同一个站点。
	CoveragePolicy string `json:"coverage_policy"`

	// 区域分布层级，只支持site。 - site:站点级别。  约束： coverage_sites中'site'字段仅支持使用ID(站点ID，通过“查询边缘站点列表”获取)，不支持name。
	CoverageLevel string `json:"coverage_level"`

	// 区域及购买数量列表。
	CoverageSites []CoverageSiteInstance `json:"coverage_sites"`
}

func (o CoverageInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoverageInstance struct{}"
	}

	return strings.Join([]string{"CoverageInstance", string(data)}, " ")
}
