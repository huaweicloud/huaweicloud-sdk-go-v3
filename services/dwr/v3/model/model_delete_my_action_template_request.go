package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMyActionTemplateRequest Request Object
type DeleteMyActionTemplateRequest struct {

	// 三方算子模板名。
	TemplateName string `json:"template_name"`
}

func (o DeleteMyActionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMyActionTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteMyActionTemplateRequest", string(data)}, " ")
}
