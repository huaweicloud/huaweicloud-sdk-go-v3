package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 作业状态。 作业的状态如下： job_init：草稿 job_submitting：提交中 job_submit_fail：提交失败 job_running：运行中（开始计费，提交作业后，返回正常结果） job_running_exception：运行异常（停止计费。作业发生运行时异常，停止运行作业） job_downloading：下载中 job_idle：空闲 job_canceling：停止中 job_cancel_success：已停止 job_cancel_fail：停止失败 job_savepointing：保存点创建中 job_arrearage_stopped：因欠费被停止（结束计费。用户账户欠费，作业停止） job_arrearage_recovering：欠费作业恢复中（用户账户欠费，账户充值，作业恢复中） job_finish：已完成
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
