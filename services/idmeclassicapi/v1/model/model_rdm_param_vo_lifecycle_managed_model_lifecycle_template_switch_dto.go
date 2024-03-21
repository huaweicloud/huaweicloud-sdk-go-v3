package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoLifecycleManagedModelLifecycleTemplateSwitchDto struct {
	Params *LifecycleManagedModelLifecycleTemplateSwitchDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoLifecycleManagedModelLifecycleTemplateSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoLifecycleManagedModelLifecycleTemplateSwitchDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoLifecycleManagedModelLifecycleTemplateSwitchDto", string(data)}, " ")
}
