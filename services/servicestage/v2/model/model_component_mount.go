package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentMount struct {

	// 挂载路径
	Path *string `json:"path,omitempty"`

	// 挂载路径的子路径
	SubPath *string `json:"sub_path,omitempty"`

	// 是否只读
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (o ComponentMount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentMount struct{}"
	}

	return strings.Join([]string{"ComponentMount", string(data)}, " ")
}
