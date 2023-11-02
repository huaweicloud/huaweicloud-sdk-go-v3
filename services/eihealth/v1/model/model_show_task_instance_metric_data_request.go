package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInstanceMetricDataRequest Request Object
type ShowTaskInstanceMetricDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	// 子任务名称
	TaskName string `json:"task_name"`

	// 子任务的并发序号
	TaskIndex *string `json:"task_index,omitempty"`

	// 子任务实例名称
	InstanceName string `json:"instance_name"`

	// 查询监控数据起始时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	FromTime *int64 `json:"from_time,omitempty"`

	// 查询监控数据截止时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	ToTime *int64 `json:"to_time,omitempty"`

	// 统计方法。枚举值，取值范围：maximum（最大值）、minimum（最小值）、average（平均值），不填时默认为maximum
	Method *string `json:"method,omitempty"`

	// 查询的监控指标名称
	MetricName string `json:"metric_name"`
}

func (o ShowTaskInstanceMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstanceMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskInstanceMetricDataRequest", string(data)}, " ")
}
