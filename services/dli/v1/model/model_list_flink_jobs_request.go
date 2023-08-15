package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobsRequest Request Object
type ListFlinkJobsRequest struct {

	// 作业类型
	JobType *string `json:"job_type,omitempty"`

	// 返回的数据条数。默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 作业名称。长度限制：0-57个字符。
	Name *string `json:"name,omitempty"`

	// 作业偏移量。
	Offset *int64 `json:"offset,omitempty"`

	// 查询结果排序，升序asc和降序desc两种可选，默认降序。
	Order *string `json:"order,omitempty"`

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 边缘父作业ID, 用于查询指定边缘作业的子作业。不带该参数时, 查询所有非边缘作业和边缘父作业, 不包括边缘子作业。
	RootJobId *int64 `json:"root_job_id,omitempty"`

	// 是否返回作业详情信息。默认为false。
	ShowDetail *bool `json:"show_detail,omitempty"`

	// 作业状态码。
	Status *string `json:"status,omitempty"`

	SysEnterpriseProjectName *string `json:"sys_enterprise_project_name,omitempty"`

	Tags *string `json:"tags,omitempty"`

	// 用户名，可作为筛选条件
	UserName *string `json:"user_name,omitempty"`
}

func (o ListFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsRequest", string(data)}, " ")
}
