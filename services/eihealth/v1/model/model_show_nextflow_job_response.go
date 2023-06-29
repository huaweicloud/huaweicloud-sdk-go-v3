package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowJobResponse Response Object
type ShowNextflowJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业的名称
	Name *string `json:"name,omitempty"`

	// 作业的描述
	Description *string `json:"description,omitempty"`

	// 作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 作业运行状态
	Status *string `json:"status,omitempty"`

	// 是否包含已被忽略的失败tasks
	HasIgnoreFailedTasks *bool `json:"has_ignore_failed_tasks,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业运行失败描述信息，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// 作业运行失败原因，当作业执行失败时会返回
	FailedReason *string `json:"failed_reason,omitempty"`

	// 流程名称
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 流程id
	WorkflowId *string `json:"workflow_id,omitempty"`

	// nextflow执行命令
	CommandLine *string `json:"command_line,omitempty"`

	// 作业参数列表
	Params *[]NextflowParamsDto `json:"params,omitempty"`

	// nextflow流程返回的全量参数列表
	NextflowParameters map[string]interface{} `json:"nextflow_parameters,omitempty"`

	// 作业标签
	ConfigFiles *[]string `json:"config_files,omitempty"`

	// nextflow config文件内容
	ConfigContext  *string `json:"config_context,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowJobResponse", string(data)}, " ")
}
