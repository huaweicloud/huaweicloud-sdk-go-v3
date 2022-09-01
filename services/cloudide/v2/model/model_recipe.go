package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Recipe struct {

	// 镜像内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 镜像类型
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o Recipe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Recipe struct{}"
	}

	return strings.Join([]string{"Recipe", string(data)}, " ")
}
