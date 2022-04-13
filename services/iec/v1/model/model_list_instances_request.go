package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 页码。 当前页面数，默认为0。 取值大于等于0，取值为0时返回第1页

	Offset *int32 `json:"offset,omitempty"`
	// 查询返回边缘实例列表当前页面的数量。 每页默认值是25，最多返回1000台边缘实例的信息，如果数据量过大建议设置成100。

	Limit *int32 `json:"limit,omitempty"`
	// 边缘实例的状态。 取值范围：ACTIVE、BUILD、DELETED、ERROR、HARD_REBOOT、MIGRATING、PAUSED、REBOOT、REBUILD、RESIZE、REVERT_RESIZE、SHUTOFF、SHELVED、SHELVED_OFFLOADED、SOFT_DELETED、SUSPENDED、VERIFY_RESIZE。  非上面范围的status字段将返回空列表。 > 当边缘实例处于中间状态时，查询范围如下： - ACTIVE，查询范围：ACTIVE，REBOOT，HARD_REBOOT，REBUILD，MIGRATING - SHUTOFF，查询范围：SHUTOFF，RESIZE，REBUILD - ERROR，查询范围：ERROR，REBUILD - VERIFY_RESIZE，查询范围：VERIFY_RESIZE，REVERT_RESIZE

	Status *string `json:"status,omitempty"`
	// 查询条件，边缘实例名称。

	Name *string `json:"name,omitempty"`
	// 边缘实例所在大区。

	Area *string `json:"area,omitempty"`
	// 边缘实例所在省份。

	Province *string `json:"province,omitempty"`
	// 边缘实例所在城市。

	City *string `json:"city,omitempty"`
	// 边缘业务ID。

	EdgecloudId *string `json:"edgecloud_id,omitempty"`
	// 站点ID。

	SiteId *string `json:"site_id,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
