package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksDatabaseUserPSinfo 修改数据库账号权限
type StarRocksDatabaseUserPSinfo struct {

	// 数据库账号名。
	UserName string `json:"user_name"`

	// 数据库列表。
	Databases *[]string `json:"databases,omitempty"`

	// DML权限。 取值范围： - 0：读写权限 - 1：只读权限 - 2：只读和设置权限 - 3：读写和设置权限
	Dml *int32 `json:"dml,omitempty"`

	// DDL权限。 取值范围： - 0：无DDL权限 - 1：有DDL权限
	Ddl *int32 `json:"ddl,omitempty"`
}

func (o StarRocksDatabaseUserPSinfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksDatabaseUserPSinfo struct{}"
	}

	return strings.Join([]string{"StarRocksDatabaseUserPSinfo", string(data)}, " ")
}
