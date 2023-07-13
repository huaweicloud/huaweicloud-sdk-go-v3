package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeDto struct {

	// **参数说明**：卷名称。
	Name string `json:"name"`

	// **参数说明**：挂载类型。
	Type *string `json:"type,omitempty"`

	// **参数说明**：源路径。
	Source string `json:"source"`

	// **参数说明**：卷挂载路径。
	Destination string `json:"destination"`

	// **参数说明**：只读，默认只读。
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (o VolumeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDto struct{}"
	}

	return strings.Join([]string{"VolumeDto", string(data)}, " ")
}
