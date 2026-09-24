package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedisPitrRestoreTimeRequest Request Object
type ListRedisPitrRestoreTimeRequest struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 查询可恢复时间点的开始时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释：** 查询可恢复时间点的结束时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释：** 偏移量，表示查询该偏移量后面的记录量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询返回记录的数量上限值。 **约束限制：** 不涉及。 **取值范围：** 1~300。 **默认取值：** 300。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRedisPitrRestoreTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedisPitrRestoreTimeRequest struct{}"
	}

	return strings.Join([]string{"ListRedisPitrRestoreTimeRequest", string(data)}, " ")
}
