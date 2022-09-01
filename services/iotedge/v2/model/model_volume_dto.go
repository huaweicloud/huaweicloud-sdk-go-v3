package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeDto struct {

	// 卷名称
	Name string `json:"name" xml:"name"`

	// 挂载类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 源路径
	Source string `json:"source" xml:"source"`

	// 卷挂载路径
	Destination string `json:"destination" xml:"destination"`

	// 只读，默认只读
	ReadOnly *bool `json:"read_only,omitempty" xml:"read_only"`
}

func (o VolumeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDto struct{}"
	}

	return strings.Join([]string{"VolumeDto", string(data)}, " ")
}
