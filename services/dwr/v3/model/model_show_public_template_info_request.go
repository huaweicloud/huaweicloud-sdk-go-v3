package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPublicTemplateInfoRequest struct {

	// 模板名称。
	TemplateName string `json:"template_name"`
}

func (o ShowPublicTemplateInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicTemplateInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicTemplateInfoRequest", string(data)}, " ")
}
