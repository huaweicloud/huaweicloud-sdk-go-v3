package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipeItem struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// UUID
	WorkspaceId string `json:"workspace_id"`

	// UUID
	PipeId string `json:"pipe_id"`

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 管道别名
	PipeAlias string `json:"pipe_alias"`

	Category *PipeCategory `json:"category"`

	// directory 目录分组
	Directory string `json:"directory"`

	// 管道描述
	Description string `json:"description"`

	ProcessStatus *PipeProcessStatus `json:"process_status"`

	ProcessError *PipeProcessError `json:"process_error"`

	OwnerType *PipeOwnerType `json:"owner_type,omitempty"`

	// 毫秒时间戳
	CreateTime int64 `json:"create_time"`

	// 毫秒时间戳
	UpdateTime int64 `json:"update_time"`

	// 毫秒时间戳
	DeleteTime int64 `json:"delete_time"`
}

func (o PipeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeItem struct{}"
	}

	return strings.Join([]string{"PipeItem", string(data)}, " ")
}
