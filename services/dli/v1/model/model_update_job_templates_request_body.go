package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateJobTemplatesRequestBody struct {

	// 名字
	Name string `json:"name"`

	// 模板内容
	Body string `json:"body"`

	// 分组
	Group *string `json:"group,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateJobTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateJobTemplatesRequestBody", string(data)}, " ")
}
