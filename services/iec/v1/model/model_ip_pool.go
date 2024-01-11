package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpPool IP池对象。 支持IPv4和IPv6
type IpPool struct {

	// 线路的ID。
	Id *string `json:"id,omitempty"`

	// 线路所属站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 线路标识。
	PoolId *string `json:"pool_id,omitempty"`

	// IPv4或IPv6线路。  取值范围： - 4：IPv4线路 - 6：IPv6线路
	IpVersion *string `json:"ip_version,omitempty"`

	Operator *Operator `json:"operator,omitempty"`

	// 线路的显示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 功能说明：表示此线路可以使用的带宽类型列表，如果列表为空，则表示该线路不能使用任何带宽  约束：线路下的ip只能加入到带宽类型在allow_share_bandwidth_types中带宽
	AllowShareBandwidthTypes *[]string `json:"allow_share_bandwidth_types,omitempty"`
}

func (o IpPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpPool struct{}"
	}

	return strings.Join([]string{"IpPool", string(data)}, " ")
}
