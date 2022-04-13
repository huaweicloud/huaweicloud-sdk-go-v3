package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标信息
type MetricInfoList struct {
	// 指标维度

	Dimensions []MetricsDimension `json:"dimensions"`
	// 指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	MetricName string `json:"metric_name"`
	// 指标命名空间，例如弹性云服务器命名空间SYS.ECS；格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空；各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace string `json:"namespace"`
	// 指标单位。

	Unit string `json:"unit"`
}

func (o MetricInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfoList struct{}"
	}

	return strings.Join([]string{"MetricInfoList", string(data)}, " ")
}
