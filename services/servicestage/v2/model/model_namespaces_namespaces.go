package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NamespacesNamespaces struct {

	// 命名空间ID。
	Id string `json:"id"`

	// 命名空间名称。
	Name string `json:"name"`
}

func (o NamespacesNamespaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespacesNamespaces struct{}"
	}

	return strings.Join([]string{"NamespacesNamespaces", string(data)}, " ")
}
