package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsDbRequestDatabases struct {

	// rds数据库id，可在查询rds数据库列表接口的ID字段获取。
	Id string `json:"id"`

	// 数据库类型 - MYSQL - ORACLE - POSTGRESQL - SQLSERVER - DAMENG - TAURUS - DWS - KINGBASE - MARIADB - GAUSSDBOPENGAUSS
	Type string `json:"type"`
}

func (o RdsDbRequestDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsDbRequestDatabases struct{}"
	}

	return strings.Join([]string{"RdsDbRequestDatabases", string(data)}, " ")
}
