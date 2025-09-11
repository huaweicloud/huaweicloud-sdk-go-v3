package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CollectWdrSnapshotRequestBody struct {

	// **参数解释**： 快照开始时间。 **约束限制**： 格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**： 快照结束时间。 **约束限制**： 格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**： 采集类型，包括实例级和组件级。 **约束限制**： 实例级则需要输入\"cluster\",组件级则需要输入组件ID。实例级和组件级无法同时输入。
	WdrType []string `json:"wdr_type"`
}

func (o CollectWdrSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectWdrSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"CollectWdrSnapshotRequestBody", string(data)}, " ")
}
