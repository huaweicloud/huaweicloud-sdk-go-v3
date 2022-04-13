package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建规则中的监控指标信息
type MetricInfoForAlarm struct {
	// 服务指标命名空间，格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空；如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace string `json:"namespace"`
	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	MetricName string `json:"metric_name"`
	// 指标维度，目前最大可添加4个维度。

	Dimensions []MetricsDimension `json:"dimensions"`
	// 创建告警规则时选择的资源分组ID，如：rg1603786526428bWbVmk4rP；如果根据资源分组创建告警规则，则resource_group_id不能为空，且alarm_type值为RESOURCE_GROUP

	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// 创建告警规则时选择的资源分组名称，如：Resource-Group-ECS-01

	ResourceGroupName *string `json:"resource_group_name,omitempty"`
}

func (o MetricInfoForAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfoForAlarm struct{}"
	}

	return strings.Join([]string{"MetricInfoForAlarm", string(data)}, " ")
}
