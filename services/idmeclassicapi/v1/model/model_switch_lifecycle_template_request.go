package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchLifecycleTemplateRequest Request Object
type SwitchLifecycleTemplateRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoLifecycleManagedModelLifecycleTemplateSwitchDto `json:"body,omitempty"`
}

func (o SwitchLifecycleTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchLifecycleTemplateRequest struct{}"
	}

	return strings.Join([]string{"SwitchLifecycleTemplateRequest", string(data)}, " ")
}
