package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStarRocksDatabaseUsersUserDetails struct {

	// 数据库账户名。
	UserName string `json:"user_name"`

	// 已授权数据库。
	DataBases []string `json:"data_bases"`

	// DML授权。 - 0：读写权限 - 1：只读权限 - 2：只读和设置权限 - 3：读写和设置权限
	Dml int32 `json:"dml"`

	// DDL授权。 - 0：无DDL权限 - 1：有DDL权限
	Ddl int32 `json:"ddl"`
}

func (o ShowStarRocksDatabaseUsersUserDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarRocksDatabaseUsersUserDetails struct{}"
	}

	return strings.Join([]string{"ShowStarRocksDatabaseUsersUserDetails", string(data)}, " ")
}
