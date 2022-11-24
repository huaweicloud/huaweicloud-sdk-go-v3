package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMyActionTemplateRequest struct {

	// 模板名称。名称必须以third开头，只能由字母、数字、下划线和中划线组成，长度小于等于32个字符。
	TemplateName string `json:"template_name"`

	Body *ThirdTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateMyActionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMyActionTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateMyActionTemplateRequest", string(data)}, " ")
}
