package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentModify struct {
	// 应用组件名称。

	Name *string `json:"name,omitempty"`
	// 描述。

	Description *string `json:"description,omitempty"`

	Source *SourceObject `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`
}

func (o ComponentModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentModify struct{}"
	}

	return strings.Join([]string{"ComponentModify", string(data)}, " ")
}
