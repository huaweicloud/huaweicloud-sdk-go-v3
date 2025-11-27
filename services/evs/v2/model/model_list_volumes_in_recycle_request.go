package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesInRecycleRequest Request Object
type ListVolumesInRecycleRequest struct {

	// **参数解释** 云硬盘名称。 可通过[查询所有云硬盘详情](evs_04_2006.xml)获取云硬盘名称。 **约束限制** 最大支持64个字符。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 云硬盘状态，取值可参考：\"[云硬盘状态](evs_04_0040.xml)\"。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释** 返回结果个数限制。 **约束限制** 不涉及。 **取值范围** 1-1000 **默认取值** 1000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 可用区信息。 可通过接口[查询所有的可用分区信息](evs_04_2081.xml)获取，也可参考[地区和终端节点](https://console.huaweicloud.com/apiexplorer/#/endpoint)获取。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释** 服务类型。 **约束限制** 不涉及。 **取值范围** - EVS：云硬盘。 - DSS：专属分布式存储服务。 **默认取值** 不涉及。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释** 分页查询时的偏移量。 **约束限制** 不涉及。 **取值范围** 取值为一个大于0小于磁盘总个数的整数。 **默认取值** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListVolumesInRecycleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesInRecycleRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesInRecycleRequest", string(data)}, " ")
}
