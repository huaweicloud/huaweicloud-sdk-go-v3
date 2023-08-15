package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMyActionTemplateToDeprecatedRequest Request Object
type UpdateMyActionTemplateToDeprecatedRequest struct {

	// 申请禁用的三方算子名称。
	TemplateName string `json:"template_name"`
}

func (o UpdateMyActionTemplateToDeprecatedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMyActionTemplateToDeprecatedRequest struct{}"
	}

	return strings.Join([]string{"UpdateMyActionTemplateToDeprecatedRequest", string(data)}, " ")
}
