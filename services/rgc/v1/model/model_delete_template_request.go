package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateRequest Request Object
type DeleteTemplateRequest struct {

	// 模板名称。
	TemplateName string `json:"template_name"`
}

func (o DeleteTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRequest", string(data)}, " ")
}
