package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductMetric struct {

	// 按云产品维度屏蔽时的指标维度信息，有多个时用\",\"连接
	DimensionName string `json:"dimension_name"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName string `json:"metric_name"`
}

func (o ProductMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductMetric struct{}"
	}

	return strings.Join([]string{"ProductMetric", string(data)}, " ")
}
