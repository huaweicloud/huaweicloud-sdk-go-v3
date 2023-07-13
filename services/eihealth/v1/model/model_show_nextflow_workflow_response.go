package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowWorkflowResponse Response Object
type ShowNextflowWorkflowResponse struct {

	// 流程id
	Id *string `json:"id,omitempty"`

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 流程描述
	Description *string `json:"description,omitempty"`

	// 流程标签
	Labels *[]string `json:"labels,omitempty"`

	// 流程的文件名
	WorkflowFile *string `json:"workflow_file,omitempty"`

	// 流程的文件名下载地址
	WorkflowFileUrl *string `json:"workflow_file_url,omitempty"`

	// 主文件名
	MainFile *string `json:"main_file,omitempty"`

	// 用户上传时使用的参数文件名
	ParamsFile *string `json:"params_file,omitempty"`

	// 流程参数列表
	Params *[]NextflowParamsDto `json:"params,omitempty"`

	// 流程的创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 流程的更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowNextflowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowWorkflowResponse", string(data)}, " ")
}
