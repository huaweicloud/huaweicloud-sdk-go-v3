package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricNameSchema **参数解释**： 指标ID，例如弹性云服务器的监控指标CPU使用率，对应的metric_name为cpu_util。各服务的命名空间可查看：[服务命名空间](ces_03_0059.xml)。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。字符串长度为[1,96]。 **默认取值**： 不涉及。
type MetricNameSchema struct {
}

func (o MetricNameSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricNameSchema struct{}"
	}

	return strings.Join([]string{"MetricNameSchema", string(data)}, " ")
}
