package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsDbListResponseDatabases struct {

	// ID
	Id string `json:"id"`

	// 数据库名称
	DbName string `json:"db_name"`

	// 实例状态。 - BUILD：表示实例正在创建。 - ACTIVE：表示实例正常。 - FAILED：表示实例异常。 - FROZEN：表示实例冻结。 - MODIFYING：表示实例正在扩容。 - REBOOTING：表示实例正在重启。 - RESTORING：表示实例正在恢复。 - MODIFYING INSTANCE TYPE：表示实例正在转主备。 - SWITCHOVER：表示实例正在主备切换。 - MIGRATING：表示实例正在迁移。 - BACKING UP：表示实例正在进行备份。 - MODIFYING DATABASE PORT：表示实例正在修改数据库端口。 - STORAGE FULL：表示实例磁盘空间满。
	Status string `json:"status"`

	// 数据库端口
	Port string `json:"port"`

	// 数据库IP
	Ip string `json:"ip"`

	// rds实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据库类型 - MYSQL - ORACLE - POSTGRESQL - SQLSERVER - DAMENG - TAURUS - DWS - KINGBASE - MARIADB - GAUSSDBOPENGAUSS
	Type string `json:"type"`

	// 版本
	Version string `json:"version"`

	// 是否支持免agent审计
	IsSupported bool `json:"is_supported"`

	// 企业项目ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o RdsDbListResponseDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsDbListResponseDatabases struct{}"
	}

	return strings.Join([]string{"RdsDbListResponseDatabases", string(data)}, " ")
}
