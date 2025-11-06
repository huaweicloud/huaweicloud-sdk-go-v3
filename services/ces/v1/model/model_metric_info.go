package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricInfo **参数解释**： 指标信息 **约束限制**： 不涉及
type MetricInfo struct {

	// 服务指标命名空间，格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空；如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。
	Namespace string `json:"namespace"`

	// **参数解释**： 指标ID，例如弹性云服务器的监控指标CPU使用率，对应的metric_name为cpu_util。各服务的命名空间可查看：[服务命名空间](ces_03_0059.xml)。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。字符串长度为[1,96]。 **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`

	// **参数解释** 指标维度，目前最大可添加4个维度。
	Dimensions []MetricsDimension `json:"dimensions"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
