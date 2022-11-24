package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowThirdTemplateInfoRequest struct {

	// 三方算子模板名称。
	TemplateName string `json:"template_name"`
}

func (o ShowThirdTemplateInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowThirdTemplateInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowThirdTemplateInfoRequest", string(data)}, " ")
}
