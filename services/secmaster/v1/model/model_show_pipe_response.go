package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipeResponse Response Object
type ShowPipeResponse struct {

	// 资源类型，例如 STREAMING_TO_INDEX
	Category *string `json:"category,omitempty"`

	// 创建者信息
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 删除时间戳
	DeleteTime *int32 `json:"delete_time,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 目录路径
	Directory *string `json:"directory,omitempty"`

	// 所有者类型，例如 USER
	OwnerType *string `json:"owner_type,omitempty"`

	// 管道别名
	PipeAlias *string `json:"pipe_alias,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 处理错误状态，例如 NONE
	ProcessError *string `json:"process_error,omitempty"`

	// 处理状态，例如 COMPLETED
	ProcessStatus *string `json:"process_status,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 资源列表
	Resources *[]IsapResource `json:"resources,omitempty"`

	Schema *Schema `json:"schema,omitempty"`

	StorageSetting *StorageSetting `json:"storage_setting,omitempty"`

	// 更新者信息
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 工作空间ID
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeResponse struct{}"
	}

	return strings.Join([]string{"ShowPipeResponse", string(data)}, " ")
}
