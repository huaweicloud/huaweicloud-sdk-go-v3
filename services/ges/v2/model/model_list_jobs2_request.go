package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobs2Request Request Object
type ListJobs2Request struct {

	// 任务结束日期，当前只支持日期，不支持时间。格式为：yyyy-MM-dd，比如2019-03-27。
	EndTime *string `json:"end_time,omitempty"`

	// 关联的图名称。
	GraphName *string `json:"graph_name,omitempty"`

	// 每页资源数量的最大值，默认为10。
	Limit *string `json:"limit,omitempty"`

	// 本次请求的起始位置，默认为0。
	Offset *string `json:"offset,omitempty"`

	// 任务开始日期，当前只支持日期，不支持时间。格式为：yyyy-MM-dd，比如2019-03-27。
	StartTime *string `json:"start_time,omitempty"`

	// 任务状态。取值为：  - running - waiting - success - failed
	Status *string `json:"status,omitempty"`
}

func (o ListJobs2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobs2Request struct{}"
	}

	return strings.Join([]string{"ListJobs2Request", string(data)}, " ")
}
