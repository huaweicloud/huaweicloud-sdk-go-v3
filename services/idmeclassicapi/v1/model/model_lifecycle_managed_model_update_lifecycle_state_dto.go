package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleManagedModelUpdateLifecycleStateDto struct {

	// 数据实例ID。
	Id string `json:"id"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`
}

func (o LifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"LifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
