package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterProtectPolicyTemplateResponse Response Object
type ShowClusterProtectPolicyTemplateResponse struct {

	// **参数解释**: 模板ID **取值范围**: 字符长度0-128
	Id *string `json:"id,omitempty"`

	// **参数解释**: 模板名称 **取值范围**: 字符长度1-255
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释**: 模板类型 **取值范围**: 字符长度1-16
	TemplateType *string `json:"template_type,omitempty"`

	// **参数解释**: 描述 **取值范围**: 字符长度0-2048
	Description *string `json:"description,omitempty"`

	// **参数解释**: 策略模板应用资源类型，多个资源类型通过分号份隔连接 **取值范围**: 字符长度1-255
	TargetKind *string `json:"target_kind,omitempty"`

	// **参数解释**: 标签 **取值范围**: 字符长度0-2048
	Tag *string `json:"tag,omitempty"`

	// **参数解释**: 推荐级别 **取值范围**: 字符长度1-5
	Level *string `json:"level,omitempty"`

	// **参数解释**: 策略模板内容 **取值范围**: 字符长度1-65535
	ConstraintTemplate *string `json:"constraint_template,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowClusterProtectPolicyTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterProtectPolicyTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterProtectPolicyTemplateResponse", string(data)}, " ")
}
