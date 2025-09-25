package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectAspRequestBody 采集asp报告请求体。
type CollectAspRequestBody struct {

	// **参数解释**: 采集开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **约束限制**: 不涉及。 **取值范围**: 当前时间前两天的时间。 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 采集结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **约束限制**: 不涉及。 **取值范围**: 当前时间前两天的时间。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`
}

func (o CollectAspRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectAspRequestBody struct{}"
	}

	return strings.Join([]string{"CollectAspRequestBody", string(data)}, " ")
}
