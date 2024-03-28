package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobProgressResponse Response Object
type ShowSqlJobProgressResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业状态。
	Status *string `json:"status,omitempty"`

	// 正在运行的子作业ID，如果作业还没开始运行或者运行结束，则子作业ID可能为空。
	SubJobId *int32 `json:"sub_job_id,omitempty"`

	// 正在运行的子作业的进度或者整个作业进度，该值只能粗略的估算子作业进度，不表示作业的详细进度。有两方面的含义： （1）如果整个作业刚开始运行或者在提交中，则进度展示为0；如果作业运行结束，则进度展示为1。此时progress表示整个作业的运行进度，因为没有子作业在运行，sub_job_id不展示。 （2）如果有子作业在运行中，则展示该子作业的运行进度，progress的计算方法为：子作业已经完成的task数除以子作业总的task数。此时progress表示子作业的运行进度，sub_job_id展示。
	Progress *float64 `json:"progress,omitempty"`

	// 正在运行作业的子作业的详细信息，一个作业可能包含多个子作业。
	SubJobs        *[]SubJob `json:"sub_jobs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSqlJobProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobProgressResponse", string(data)}, " ")
}
