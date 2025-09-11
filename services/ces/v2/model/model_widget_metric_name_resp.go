package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WidgetMetricNameResp **参数解释** 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **取值范围** 长度为[1,96]个字符
type WidgetMetricNameResp struct {
}

func (o WidgetMetricNameResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetMetricNameResp struct{}"
	}

	return strings.Join([]string{"WidgetMetricNameResp", string(data)}, " ")
}
