package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseTemplateVersion 模板版本基本信息
type BaseTemplateVersion struct {

	// 模板的唯一ID，由模板服务随机生成
	TemplateId string `json:"template_id"`

	// 模板（Template）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
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
