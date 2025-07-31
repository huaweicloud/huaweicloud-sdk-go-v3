package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterProtectPolicyTemplateResponse Response Object
type ShowClusterProtectPolicyTemplateResponse struct {

	// 模板ID
	Id *string `json:"id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板类型
	TemplateType *string `json:"template_type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 策略模板应用资源类型，多个资源类型通过分号份隔连接
	TargetKind *string `json:"target_kind,omitempty"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 推荐级别
	Level *string `json:"level,omitempty"`

	// 策略模板内容
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
