package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentItem struct {

	// 组件ID。
	Id *string `json:"id,omitempty"`

	// 组件名称。
	Name *string `json:"name,omitempty"`

	// 资源信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	Spec *ComponentSpec `json:"spec,omitempty"`
}

func (o ComponentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentItem struct{}"
	}

	return strings.Join([]string{"ComponentItem", string(data)}, " ")
}
