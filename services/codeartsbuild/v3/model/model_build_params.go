package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildParams 构建参数约束列表
type BuildParams struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 名称
	Title *string `json:"title,omitempty"`

	// 简要构建信息列表
	Params *[]Params `json:"params,omitempty"`
}

func (o BuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildParams struct{}"
	}

	return strings.Join([]string{"BuildParams", string(data)}, " ")
}
