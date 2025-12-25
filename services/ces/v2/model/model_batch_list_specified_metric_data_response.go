package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSpecifiedMetricDataResponse Response Object
type BatchListSpecifiedMetricDataResponse struct {

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**: 资源维度, 多维度逗号分隔。 **取值范围**: 必须以字母开头，只能包含0-9/a-z/A-Z/_/-/,。每个维度必须以字母开头，每个维度长度最短1，最长32，多个维度直接用,分隔。
	MetricDimension *string `json:"metric_dimension,omitempty"`

	// **参数解释** 监控数据列表
	DataPoints     *[]MetricDataPoint `json:"data_points,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchListSpecifiedMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSpecifiedMetricDataResponse struct{}"
	}

	return strings.Join([]string{"BatchListSpecifiedMetricDataResponse", string(data)}, " ")
}
