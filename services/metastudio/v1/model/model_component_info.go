package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件信息。
type ComponentInfo struct {

	// 组件名称。
	ComponentName string `json:"component_name"`

	// 组件类型。
	ComponentType string `json:"component_type"`

	// 组件描述。
	ComponentDesc *string `json:"component_desc,omitempty"`
}

func (o ComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfo struct{}"
	}

	return strings.Join([]string{"ComponentInfo", string(data)}, " ")
}
