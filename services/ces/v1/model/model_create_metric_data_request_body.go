package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetricDataRequestBody 添加一条或多条自定义指标监控数据，请求参数。
type CreateMetricDataRequestBody struct {
	Metric *MetricInfo `json:"metric"`

	// **参数解释**： 数据的有效期，超出该有效期则自动删除该数据，单位秒。 **约束限制**： 不涉及。 **取值范围**： 大小为[1,604800]的整数。 **默认取值**： 不涉及。
	Ttl int32 `json:"ttl"`

	// **参数解释**： 数据收集时间 。UNIX时间戳，单位毫秒。 **约束限制**： 不涉及。 **取值范围**： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-3天+10秒，当前时间+10分钟-10秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。 **默认取值**： 不涉及。
	CollectTime int64 `json:"collect_time"`

	// **参数解释**： 待添加的监控指标数据的值。 **约束限制**： 不涉及。 **取值范围**： 数值类型支持“整数”或“浮点数”。 **默认取值**： 不涉及。
	Value float64 `json:"value"`

	// **参数解释**： 数据的单位。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符。 **默认取值**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 数据类型。 **约束限制**： 不涉及。 **取值范围**： 枚举值，只能是\"int\"或\"float\"。 - int：整数 - float：浮点数 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o CreateMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMetricDataRequestBody", string(data)}, " ")
}
