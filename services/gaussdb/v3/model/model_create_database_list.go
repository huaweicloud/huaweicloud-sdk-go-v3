package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseList 数据库名称列表，即创建数据库用户时同步将列表中的数据库授权给用户，列表最大长度为50。列表可以为空，在需要给该用户授权某数据库时，调用数据库用户授权接口即可。
type CreateDatabaseList struct {

	// 数据库名称。
	Name string `json:"name"`

	// 是否为只读权限： - true，表示只读。 - false，表示可读写。
	Readonly bool `json:"readonly"`
}

func (o CreateDatabaseList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseList struct{}"
	}

	return strings.Join([]string{"CreateDatabaseList", string(data)}, " ")
}
