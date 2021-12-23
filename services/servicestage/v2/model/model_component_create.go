package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentCreate struct {
	// 应用组件名称。

	Name string `json:"name"`

	Runtime *RuntimeType `json:"runtime"`

	Category *ComponentCategory `json:"category"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty"`
	// 描述。

	Description *string `json:"description,omitempty"`

	Source *SourceObject `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`
}

func (o ComponentCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentCreate struct{}"
	}

	return strings.Join([]string{"ComponentCreate", string(data)}, " ")
}
