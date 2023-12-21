package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Distribution 实例分布对象
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

	// 创建边缘实例是否开启IPv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 创建IPv6边缘实例是否支持公网访问。
	Ipv6BandwidthEnable *bool `json:"ipv6_bandwidth_enable,omitempty"`

	// IPv6线路ID。IPv6场景下，使用该线路下的子网分配IPv6端口。
	PoolIdV6 *string `json:"pool_id_v6,omitempty"`
}

func (o Distribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Distribution struct{}"
	}

	return strings.Join([]string{"Distribution", string(data)}, " ")
}
