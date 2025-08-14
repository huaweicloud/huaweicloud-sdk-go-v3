package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateMetadataResponse Response Object
type ShowTemplateMetadataResponse struct {

	// 模板的唯一ID，由模板服务随机生成
	TemplateId string `json:"template_id"`

	// 模板（Template）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	TemplateName string `json:"template_name"`

	// 模板的描述。可用于客户识别自己的模板
	TemplateDescription *string `json:"template_description,omitempty"`

	// 模板的生成时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime string `json:"create_time"`

	// 模板的更新时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	UpdateTime     string `json:"update_time"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTemplateMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateMetadataResponse", string(data)}, " ")
}
