package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateRequest Request Object
type ShowTemplateRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 模板id
	TemplateId string `json:"template_id"`
}

func (o ShowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}
