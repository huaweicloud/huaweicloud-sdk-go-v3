package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentMount struct {

	// 挂载路径
	Path *string `json:"path,omitempty"`

	// 挂载路径的子路径
	SubPath *string `json:"subPath,omitempty"`

	// 是否只读
	ReadOnly *bool `json:"readOnly,omitempty"`
}

func (o ComponentMount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentMount struct{}"
	}

	return strings.Join([]string{"ComponentMount", string(data)}, " ")
}
