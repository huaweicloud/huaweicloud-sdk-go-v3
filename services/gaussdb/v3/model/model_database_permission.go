package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabasePermission 数据库权限列表，列表最大长度为50。
type DatabasePermission struct {

	// 数据库名称。
	Name string `json:"name"`

	// 是否为只读权限： - true，表示只读。 - false，表示可读写。
	Readonly bool `json:"readonly"`
}

func (o DatabasePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabasePermission struct{}"
	}

	return strings.Join([]string{"DatabasePermission", string(data)}, " ")
}
