package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleManagedModelLifecycleTemplateSwitchDto struct {

	// 唯一标识。
	Id string `json:"id"`

	LifecycleTemplate *ObjectReferenceParamDto `json:"lifecycleTemplate"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState"`
}

func (o LifecycleManagedModelLifecycleTemplateSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleManagedModelLifecycleTemplateSwitchDto struct{}"
	}

	return strings.Join([]string{"LifecycleManagedModelLifecycleTemplateSwitchDto", string(data)}, " ")
}
