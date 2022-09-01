package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentCreate struct {

	// 应用组件名称。
	Name string `json:"name" xml:"name"`

	Runtime *RuntimeType `json:"runtime" xml:"runtime"`

	Category *ComponentCategory `json:"category" xml:"category"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty" xml:"sub_category"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	Source *SourceObject `json:"source,omitempty" xml:"source"`

	Build *Build `json:"build,omitempty" xml:"build"`
}

func (o ComponentCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentCreate struct{}"
	}

	return strings.Join([]string{"ComponentCreate", string(data)}, " ")
}
