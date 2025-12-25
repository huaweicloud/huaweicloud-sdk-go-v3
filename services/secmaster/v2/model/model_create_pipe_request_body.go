package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePipeRequestBody struct {

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 管道别名
	PipeAlias string `json:"pipe_alias"`

	Category *PipeCategory `json:"category"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// 管道描述
	Description *string `json:"description,omitempty"`

	Schema *PipeSchema `json:"schema"`

	StorageSetting *PipeStorageSetting `json:"storage_setting"`

	DisplaySetting *TableDisplaySetting `json:"display_setting,omitempty"`
}

func (o CreatePipeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePipeRequestBody", string(data)}, " ")
}
