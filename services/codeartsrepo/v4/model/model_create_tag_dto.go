package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTagDto struct {

	// 新建标签名称
	TagName string `json:"tag_name"`

	// 基于ref（分支、commitid）创建
	Ref string `json:"ref"`

	// 标签描述信息
	Message *string `json:"message,omitempty"`
}

func (o CreateTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagDto struct{}"
	}

	return strings.Join([]string{"CreateTagDto", string(data)}, " ")
}
