package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateRequestBody CreateTemplate API的响应体
type CreateTemplateRequestBody struct {

	// 模板版本的描述。可用于客户识别自己的模板版本
	VersionDescription *string `json:"version_description,omitempty"`

	// HCL模板，描述了模板中使用的资源 template_body 和 template_uri 有且仅有一个存在
	TemplateBody *string `json:"template_body,omitempty"`

	// HCL模板的obs链接，该模板描述了资源的目标状态  对应的文件应该是纯tf文件或zip压缩包  纯tf文件需要以`.tf`或者`.tf.json`结尾，并遵守hcl语法  压缩包目前只支持zip格式，文件需要以\".zip\"结尾。解压后的文件不得包含\".tfvars\"文件  template_body 和 template_uri 有且仅有一个存在
	TemplateUri *string `json:"template_uri,omitempty"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	// 模板的描述。可用于客户识别自己的模板
	TemplateDescription *string `json:"template_description,omitempty"`
}

func (o CreateTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequestBody", string(data)}, " ")
}
