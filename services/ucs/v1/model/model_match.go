package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Match struct {

	// 生效类型，当前预置，填写不会生效
	Kinds *[]ResourceKinds `json:"kinds,omitempty"`

	// 生效的命名空间
	Namespaces *[]string `json:"namespaces,omitempty"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
