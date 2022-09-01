package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentAmb struct {

	// 组件ID。
	ComponentId *string `json:"componentId,omitempty" xml:"componentId"`

	// 组件名称。
	ComponentName *string `json:"componentName,omitempty" xml:"componentName"`

	// 组件版本。
	ComponentVersion *string `json:"componentVersion,omitempty" xml:"componentVersion"`

	// 组件描述信息。
	ComponentDesc *string `json:"componentDesc,omitempty" xml:"componentDesc"`
}

func (o ComponentAmb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentAmb struct{}"
	}

	return strings.Join([]string{"ComponentAmb", string(data)}, " ")
}
