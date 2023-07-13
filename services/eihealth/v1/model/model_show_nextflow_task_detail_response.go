package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowTaskDetailResponse Response Object
type ShowNextflowTaskDetailResponse struct {

	// task id
	Id *string `json:"id,omitempty"`

	// task执行命令
	Command *string `json:"command,omitempty"`

	// task状态
	Status *string `json:"status,omitempty"`

	// task失败后的策略
	ErrorAction *string `json:"error_action,omitempty"`

	// task退出状态码
	Exit *int32 `json:"exit,omitempty"`

	// task执行路径
	WorkDir *string `json:"work_dir,omitempty"`

	// task执行的环境变量值
	Environment *string `json:"environment,omitempty"`

	// 子任务运行环境列表
	Module *[]string `json:"module,omitempty"`

	// 容器名称
	Container *string `json:"container,omitempty"`

	// 执行次数
	Attempt *int32 `json:"attempt,omitempty"`

	// 临时工作目录
	Scratch *string `json:"scratch,omitempty"`

	ExecutionTime *NextflowTaskExecutionTime `json:"execution_time,omitempty"`

	ResourceRequested *NextflowTaskResourceRequested `json:"resource_requested,omitempty"`

	ResourceUsage  *NextflowTaskResourceUsage `json:"resource_usage,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowNextflowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowTaskDetailResponse", string(data)}, " ")
}
