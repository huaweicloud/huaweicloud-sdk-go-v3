package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowClickHouseDatabaseUsersUserDetails struct {

	// 数据库账户名。
	UserName string `json:"user_name"`

	// 已授权数据库。
	Databases []string `json:"databases"`

	// DML权限。 取值范围： - 1：只读权限 - 2：读取和设置权限
	Dml int32 `json:"dml"`

	// DDL授权。 取值范围： - 0：无DDL权限 - 1：有DDL权限
	Ddl int32 `json:"ddl"`
}

func (o ShowClickHouseDatabaseUsersUserDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseDatabaseUsersUserDetails struct{}"
	}

	return strings.Join([]string{"ShowClickHouseDatabaseUsersUserDetails", string(data)}, " ")
}
