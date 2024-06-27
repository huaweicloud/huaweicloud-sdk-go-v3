package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseDatabaseUserPSinfo 修改数据库账号权限。
type ClickHouseDatabaseUserPSinfo struct {

	// 数据库账号名。
	UserName string `json:"user_name"`

	// 数据库列表。“*”表示所有数据库。
	Databases []string `json:"databases"`

	// DML权限。 取值范围： - 1：只读权限 - 2：读取和设置权限
	Dml int32 `json:"dml"`
}

func (o ClickHouseDatabaseUserPSinfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseDatabaseUserPSinfo struct{}"
	}

	return strings.Join([]string{"ClickHouseDatabaseUserPSinfo", string(data)}, " ")
}
