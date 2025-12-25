package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePipeRequestBody struct {

	// 管道别名
	PipeAlias *string `json:"pipe_alias,omitempty"`

	// 管道描述
	Description *string `json:"description,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	Category *PipeCategory `json:"category,omitempty"`
}

func (o UpdatePipeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePipeRequestBody", string(data)}, " ")
}
