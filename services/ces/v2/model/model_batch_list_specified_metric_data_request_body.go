package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSpecifiedMetricDataRequestBody ***参数解释*** 批量查询指标数据请求体 ***约束限制*** 不涉及
type BatchListSpecifiedMetricDataRequestBody struct {

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为64。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`

	// **参数解释**: 指标维度, 多维度逗号分隔。 **约束限制**: 不涉及。 **取值范围**: 必须以字母开头，只能包含0-9/a-z/A-Z/_/-/,。每个维度必须以字母开头，每个维度长度最短1，最长32，多个维度直接用,分隔。 **默认取值**: 不涉及。
	MetricDimension string `json:"metric_dimension"`

	// **参数解释**: 查询监控数据的开始时间，格式为时间戳, 单位毫秒。 **约束限制**: from必须小于to, to和from的时间间隔必须在5分钟内。 **取值范围**: 最小值为0。 **默认取值**: 不涉及。
	From *int64 `json:"from,omitempty"`

	// **参数解释**: 查询监控数据的结束时间，格式为时间戳, 单位毫秒。 **约束限制**: from必须小于to, to和from的时间间隔必须在5分钟内。 **取值范围**: 最小值为0。 **默认取值**: 不涉及。
	To *int64 `json:"to,omitempty"`

	// **参数解释**: 分页大小。 **约束限制**: 不涉及。 **取值范围**: 最小值为1，最大值为1000。 **默认取值**: 100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 分页偏移量。 **约束限制**: 不涉及。 **取值范围**: 最小值为0，最大值为9999999。 **默认取值**: 0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o BatchListSpecifiedMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSpecifiedMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListSpecifiedMetricDataRequestBody", string(data)}, " ")
}
