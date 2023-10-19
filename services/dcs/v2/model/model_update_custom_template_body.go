package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCustomTemplateBody struct {

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板的描述信息
	Description *string `json:"description,omitempty"`

	// 参数配置信息
	Params map[string]string `json:"params,omitempty"`
}

func (o UpdateCustomTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomTemplateBody struct{}"
	}

	return strings.Join([]string{"UpdateCustomTemplateBody", string(data)}, " ")
}
