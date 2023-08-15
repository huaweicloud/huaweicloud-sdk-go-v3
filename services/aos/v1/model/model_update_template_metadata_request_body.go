package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateMetadataRequestBody 更新模板元数据请求体
type UpdateTemplateMetadataRequestBody struct {

	// 模板的唯一ID，由模板服务随机生成
	TemplateId *string `json:"template_id,omitempty"`

	// 模板的描述。可用于客户识别自己的模板
	TemplateDescription *string `json:"template_description,omitempty"`
}

func (o UpdateTemplateMetadataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateMetadataRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTemplateMetadataRequestBody", string(data)}, " ")
}
