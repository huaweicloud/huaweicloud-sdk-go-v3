package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentModify struct {

	// 应用组件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	Source *SourceObject `json:"source,omitempty" xml:"source"`

	Build *Build `json:"build,omitempty" xml:"build"`
}

func (o ComponentModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentModify struct{}"
	}

	return strings.Join([]string{"ComponentModify", string(data)}, " ")
}
