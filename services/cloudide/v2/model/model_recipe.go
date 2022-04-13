package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Recipe struct {
	// 镜像内容

	Content *string `json:"content,omitempty"`
	// 镜像类型

	Type *string `json:"type,omitempty"`
}

func (o Recipe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Recipe struct{}"
	}

	return strings.Join([]string{"Recipe", string(data)}, " ")
}
