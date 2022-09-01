package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库权限列表，列表最大长度为50。
type DatabasePermission struct {

	// 数据库名称。
	Name string `json:"name" xml:"name"`

	// 是否为只读权限：true表示只读,false表示可读写。
	Readonly bool `json:"readonly" xml:"readonly"`
}

func (o DatabasePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabasePermission struct{}"
	}

	return strings.Join([]string{"DatabasePermission", string(data)}, " ")
}
