package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemTemplateDetailRequest Request Object
type ShowSystemTemplateDetailRequest struct {

	// 系统算子模板名称.。名称必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符，且不能重名。
	TemplateName string `json:"template_name"`
}

func (o ShowSystemTemplateDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemTemplateDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSystemTemplateDetailRequest", string(data)}, " ")
}
