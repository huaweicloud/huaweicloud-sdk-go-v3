package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// parse-template-variables request parameters
type ParseTemplateVariablesRequestBody struct {

	// HCL模板，描述了资源的目标状态 template_body 和 template_uri 有且仅有一个存在
	TemplateBody *string `json:"template_body,omitempty"`

	// HCL模板文件的uri，代码需从该uri下载HCL模板获取模板内容。其描述了资源的目标状态 template_body 和 template_uri 有且仅有一个存在
	TemplateUri *string `json:"template_uri,omitempty"`
}

func (o ParseTemplateVariablesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseTemplateVariablesRequestBody struct{}"
	}

	return strings.Join([]string{"ParseTemplateVariablesRequestBody", string(data)}, " ")
}
