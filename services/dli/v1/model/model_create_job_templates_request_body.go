package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateJobTemplatesRequestBody struct {

	// 类型。SQL：sql模板。SPARK：spark模板。
	Type string `json:"type"`

	// 名字
	Name string `json:"name"`

	// 模板内容
	Body string `json:"body"`

	// 分组
	Group *string `json:"group,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`
}

func (o CreateJobTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateJobTemplatesRequestBody", string(data)}, " ")
}
