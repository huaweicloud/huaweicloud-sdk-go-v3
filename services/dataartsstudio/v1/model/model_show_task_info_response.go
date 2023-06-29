package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInfoResponse Response Object
type ShowTaskInfoResponse struct {

	// 任务id
	Id *string `json:"id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务描述
	Description *string `json:"description,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 产品id
	ProjectId *string `json:"project_id,omitempty"`

	// 目录id
	DirId *string `json:"dir_id,omitempty"`

	ScheduleConfig *SchedulerInfo `json:"schedule_config,omitempty"`

	// 自定义元数据信息
	ParameterConfig *[]CustomMetadata `json:"parameter_config,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 最后一次执行时间
	LastRunTime *string `json:"last_run_time,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 下一次执行时间
	NextRunTime *string `json:"next_run_time,omitempty"`

	// 责任人
	DutyPerson *string `json:"duty_person,omitempty"`

	// 修改类型
	UpdateType *string `json:"update_type,omitempty"`

	// 数据来源类型
	DataSourceType *string `json:"data_source_type,omitempty"`

	// 任务信息Map<String, Object>
	TaskConfig *interface{} `json:"task_config,omitempty"`

	// 数据来源工作空间id
	DataSourceWorkspaceId *string `json:"data_source_workspace_id,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ShowTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInfoResponse", string(data)}, " ")
}
