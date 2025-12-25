package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComponentTemplateRequest Request Object
type UpdateComponentTemplateRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 插件配置模板ID
	TemplateId string `json:"template_id"`

	Body *CreateTemplateReq `json:"body,omitempty"`
}

func (o UpdateComponentTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateComponentTemplateRequest", string(data)}, " ")
}
