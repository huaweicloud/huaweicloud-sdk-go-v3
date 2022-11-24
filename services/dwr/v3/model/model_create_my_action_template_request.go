package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMyActionTemplateRequest struct {

	// 算子名称，名称必须以以third开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符，且不能重名。
	TemplateName string `json:"template_name"`

	Body *ThirdTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateMyActionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMyActionTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateMyActionTemplateRequest", string(data)}, " ")
}
