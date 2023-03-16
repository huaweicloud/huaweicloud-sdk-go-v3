package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDataJobResponse struct {

	// 已完成的数据列表
	CompleteData *[]string `json:"complete_data,omitempty"`

	// 正在执行的数据列表
	RunningData *[]string `json:"running_data,omitempty"`

	// 数据作业创建者
	Creator *string `json:"creator,omitempty"`

	// 非本项目操作场景下源项目名称
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 非本项目操作场景下源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 数据作业ID
	Id *string `json:"id,omitempty"`

	// 数据作业名称
	Name *string `json:"name,omitempty"`

	// 数据列表
	Sources *[]string `json:"sources,omitempty"`

	// 数据作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 数据作业结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 数据作业状态
	Status *string `json:"status,omitempty"`

	// 数据列表
	Destinations *[]string `json:"destinations,omitempty"`

	// 数据作业类型
	Type *string `json:"type,omitempty"`

	// 数据作业失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 附加信息
	Additions      *string `json:"additions,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataJobResponse struct{}"
	}

	return strings.Join([]string{"ShowDataJobResponse", string(data)}, " ")
}
