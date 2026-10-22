package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBiactiveInstanceCandidatesRequest Request Object
type ListBiactiveInstanceCandidatesRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 目标Region的编码。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TargetRegionCode string `json:"target_region_code"`

	// **参数解释：** 按目标实例名称过滤。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	// **参数解释：** 数据同步模式。 **约束限制：** 不涉及。 **取值范围：** - TABLE：表级数据同步模式。 **默认取值：** 不传该参数时表示实例级数据同步模式。
	DataSyncMode *string `json:"data_sync_mode,omitempty"`

	// **参数解释：** 索引位置偏移量，表示从最新一条记录开始偏移offset条数据后查询。 **约束限制：** 取值大于或等于0。 **取值范围：** 不涉及。 **默认取值：** 不传该参数时，查询偏移量默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询实例个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBiactiveInstanceCandidatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBiactiveInstanceCandidatesRequest struct{}"
	}

	return strings.Join([]string{"ListBiactiveInstanceCandidatesRequest", string(data)}, " ")
}
