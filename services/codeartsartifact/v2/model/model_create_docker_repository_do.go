package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDockerRepositoryDo struct {

	// 仓库格式
	Format string `json:"format"`

	// 仓库展示名称
	DisplayName string `json:"display_name"`

	// 仓库描述
	Description *string `json:"description,omitempty"`

	// 仓库类型
	Type string `json:"type"`
}

func (o CreateDockerRepositoryDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDockerRepositoryDo struct{}"
	}

	return strings.Join([]string{"CreateDockerRepositoryDo", string(data)}, " ")
}
