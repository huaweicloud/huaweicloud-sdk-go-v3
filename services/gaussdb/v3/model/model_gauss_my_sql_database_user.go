package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussMySqlDatabaseUser 需要授权的用户信息
type GaussMySqlDatabaseUser struct {

	// 数据库用户名
	Name string `json:"name"`

	// 主机地址
	Host string `json:"host"`

	// 是否为只读权限： - true，表示只读。 - false，表示可读写。
	Readonly bool `json:"readonly"`
}

func (o GaussMySqlDatabaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussMySqlDatabaseUser struct{}"
	}

	return strings.Join([]string{"GaussMySqlDatabaseUser", string(data)}, " ")
}
