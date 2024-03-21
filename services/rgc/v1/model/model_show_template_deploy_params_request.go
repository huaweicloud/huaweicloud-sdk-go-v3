package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateDeployParamsRequest Request Object
type ShowTemplateDeployParamsRequest struct {

	// 模板名称。
	TemplateName string `json:"template_name"`

	// 模板版本。
	Version string `json:"version"`
}

func (o ShowTemplateDeployParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateDeployParamsRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateDeployParamsRequest", string(data)}, " ")
}
