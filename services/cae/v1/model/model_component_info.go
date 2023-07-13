package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentInfo struct {

	// 组件ID。
	ComponentId *string `json:"component_id,omitempty"`

	// 组件名称。
	ComponentName *string `json:"component_name,omitempty"`
}

func (o ComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfo struct{}"
	}

	return strings.Join([]string{"ComponentInfo", string(data)}, " ")
}
