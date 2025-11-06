package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordingRuleRequest struct {

	// 预聚合规则。 待创建的预聚合规则详细信息。支持如下子参数： - groups：规则组。一份RecordingRule.yaml可以配置多组规则组。 - name：规则组名称。规则组名称必须唯一。 - interval：规则组的执行周期。默认60s。（可选） - rules：规则。一个规则组可以包含多条规则。 - record：规则的名称。聚合规则的名称必须符合  [Prometheus指标名称规范](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels)。  - expr：计算表达式。Prometheus监控将通过该表达式计算得出预聚合指标。计算表达式必须符合[PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/)。 - labels：指标的标签。标签必须符合[Prometheus指标标签规范](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels)。（可选）
	RecordingRule string `json:"recording_rule"`
}

func (o RecordingRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordingRuleRequest struct{}"
	}

	return strings.Join([]string{"RecordingRuleRequest", string(data)}, " ")
}
