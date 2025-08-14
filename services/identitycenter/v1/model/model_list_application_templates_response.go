package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationTemplatesResponse Response Object
type ListApplicationTemplatesResponse struct {

	// 应用程序模板列表
	ApplicationTemplates *[]ApplicationTemplateDto `json:"application_templates,omitempty"`
	HttpStatusCode       int                       `json:"-"`
}

func (o ListApplicationTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationTemplatesResponse", string(data)}, " ")
}
