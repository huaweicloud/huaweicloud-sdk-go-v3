package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentItem struct {

	// 组件id。
	Id *string `json:"id,omitempty"`

	// 组件名称。
	Name *string `json:"name,omitempty"`

	// 组件状态。
	Status *string `json:"status,omitempty"`

	// 资源信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	Spec *ComponentSpec `json:"spec,omitempty"`
}

func (o ComponentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentItem struct{}"
	}

	return strings.Join([]string{"ComponentItem", string(data)}, " ")
}
