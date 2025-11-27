package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceKinds struct {

	// 资源所属的API组
	ApiGroups *[]string `json:"apiGroups,omitempty"`

	// 资源类型
	Kinds *[]string `json:"kinds,omitempty"`
}

func (o ResourceKinds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceKinds struct{}"
	}

	return strings.Join([]string{"ResourceKinds", string(data)}, " ")
}
