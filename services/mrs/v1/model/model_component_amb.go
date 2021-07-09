package model

import (
	"encoding/json"

	"strings"
)

type ComponentAmb struct {
	// 组件ID。

	ComponentId *string `json:"componentId,omitempty"`
	// 组件名称。

	ComponentName *string `json:"componentName,omitempty"`
	// 组件版本。

	ComponentVersion *string `json:"componentVersion,omitempty"`
	// 组件描述信息。

	ComponentDesc *string `json:"componentDesc,omitempty"`
}

func (o ComponentAmb) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ComponentAmb struct{}"
	}

	return strings.Join([]string{"ComponentAmb", string(data)}, " ")
}
