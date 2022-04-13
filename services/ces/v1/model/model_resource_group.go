package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源分组中的资源信息
type ResourceGroup struct {
	// 资源类型。即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace *string `json:"namespace,omitempty"`
	// 一个或者多个资源维度。

	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
	// 资源分组中该资源的当前状态，值可为health、unhealth、no_alarm_rule；health表示健康，unhealth表示不健康，no_alarm_rule表示未设置告警规则。

	Status *string `json:"status,omitempty"`
}

func (o ResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroup struct{}"
	}

	return strings.Join([]string{"ResourceGroup", string(data)}, " ")
}
