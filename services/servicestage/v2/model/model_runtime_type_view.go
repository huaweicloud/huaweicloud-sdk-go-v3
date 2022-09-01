package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 运行时参数。
type RuntimeTypeView struct {

	// 类型名称。
	TypeName *string `json:"type_name,omitempty" xml:"type_name"`

	// 显示名称。
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 容器默认端口。
	ContainerDefaultPort *int32 `json:"container_default_port,omitempty" xml:"container_default_port"`

	// 类型描述。
	TypeDesc *string `json:"type_desc,omitempty" xml:"type_desc"`
}

func (o RuntimeTypeView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeTypeView struct{}"
	}

	return strings.Join([]string{"RuntimeTypeView", string(data)}, " ")
}
