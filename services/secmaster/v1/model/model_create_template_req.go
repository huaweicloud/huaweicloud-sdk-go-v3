package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateReq 创建/更新插件配置模板请求体
type CreateTemplateReq struct {

	// 插件ID
	ComponentId string `json:"component_id"`

	// 插件模板的名称
	TemplateName string `json:"template_name"`

	// 插件action的配置内容
	TaskConfig string `json:"task_config"`
}

func (o CreateTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateTemplateReq", string(data)}, " ")
}
