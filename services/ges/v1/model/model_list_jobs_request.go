package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListJobsRequest struct {
	// 任务结束日期，当前只支持日期，不支持时间。格式为：yyyy-MM-dd，比如2019-03-27。

	EndTime *string `json:"endTime,omitempty"`
	// 关联的图名称

	GraphName *string `json:"graph_name,omitempty"`
	// 每页资源数量的最大值，默认为10。

	Limit *string `json:"limit,omitempty"`
	// 本次请求的起始位置，默认为0。

	Offset *string `json:"offset,omitempty"`
	// 任务开始日期，当前只支持日期，不支持时间。格式为：yyyy-MM-dd，比如2019-03-27。

	StartTime *string `json:"startTime,omitempty"`
	// 任务状态。取值为：  - running - waiting - success - failed

	Status *string `json:"status,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
