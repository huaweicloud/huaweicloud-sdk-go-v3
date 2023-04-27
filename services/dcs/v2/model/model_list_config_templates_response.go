package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigTemplatesResponse struct {

	// 模板个数。
	TemplateNum float32 `json:"template_num,omitempty"`

	// 模板的详情数组。
	ConfigTemplates *[]ConfigTemplatesListInfo `json:"config_templates,omitempty"`
	HttpStatusCode  int                        `json:"-"`
}

func (o ListConfigTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigTemplatesResponse", string(data)}, " ")
}
