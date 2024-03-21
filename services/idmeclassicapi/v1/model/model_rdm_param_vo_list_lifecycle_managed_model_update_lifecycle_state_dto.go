package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct {

	// 参数对象。
	Params *[]LifecycleManagedModelUpdateLifecycleStateDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
