package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseTemplateVersion 模板版本基本信息
type BaseTemplateVersion struct {

	// 模板的唯一ID，由模板服务随机生成
	TemplateId string `json:"template_id"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	// 模板版本的描述。可用于客户识别自己的模板版本
	VersionDescription *string `json:"version_description,omitempty"`

	// 版本创建时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime string `json:"create_time"`
}

func (o BaseTemplateVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseTemplateVersion struct{}"
	}

	return strings.Join([]string{"BaseTemplateVersion", string(data)}, " ")
}
