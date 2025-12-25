package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipeResponse Response Object
type CreatePipeResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// UUID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// UUID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 管道别名
	PipeAlias *string `json:"pipe_alias,omitempty"`

	Category *PipeCategory `json:"category,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// 管道描述
	Description *string `json:"description,omitempty"`

	ProcessStatus *PipeProcessStatus `json:"process_status,omitempty"`

	ProcessError *PipeProcessError `json:"process_error,omitempty"`

	OwnerType *PipeOwnerType `json:"owner_type,omitempty"`

	// 管道资源
	Resources *[]PipeResource `json:"resources,omitempty"`

	Schema *PipeSchema `json:"schema,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreatePipeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeResponse struct{}"
	}

	return strings.Join([]string{"CreatePipeResponse", string(data)}, " ")
}
