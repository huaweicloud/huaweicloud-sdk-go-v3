package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataBase 数据库信息
type DataBase struct {

	// 数据库ID
	Id *string `json:"id,omitempty"`

	// 数据库名称
	Name string `json:"name"`

	// 添加的数据库类型： - MYSQL：MySQL - ORACLE：Oracle - POSTGRESQL: PostgreSQL - SQLSERVER: SqlServer - DAMENG: Dameng - TAURUS: Taurus - DWS: Dws - KINGBASE: Kingbase - GAUSSDBOPENGAUSS：GaussDB(for openGauss) - GREENPLUM：Greenplum - HIGHGO：HighGo - SHENTONG：ShenTong - GBASE8A：GBase 8a - GBASE8S：GBase 8s - GBASEXDM：GBase XDM - MONGODB：MongoDB - DDS：DDS（Document Database Service）
	Type string `json:"type"`

	// 数据库版本
	Version string `json:"version"`

	// 数据库字符集 - GBK：GBK - UTF8：UTF8
	Charset string `json:"charset"`

	// 数据库IP
	Ip string `json:"ip"`

	// 数据库端口
	Port string `json:"port"`

	// 数据库操作系统
	Os string `json:"os"`

	// 实例状态 - ON :开启 - OFF : 关闭
	Status string `json:"status"`

	// 数据库实例名
	InstanceName string `json:"instance_name"`

	// 数据库的运行状态 - ACTIVE：运行中 - SHUTOFF：已关闭 - ERROR：故障
	AuditStatus *string `json:"audit_status,omitempty"`

	// agent的唯一ID
	AgentUrl *[]string `json:"agent_url,omitempty"`

	// 数据库分类 - RDS: 表示RDS数据库 - ECS: 自建数据库
	DbClassification string `json:"db_classification"`

	// rds实例审计开关状态不匹配。当数据库审计开启且rds侧日志上传开关关闭时该字段为true。
	RdsAuditSwitchMismatch *bool `json:"rds_audit_switch_mismatch,omitempty"`

	// RDS数据库的ID。
	RdsId *string `json:"rds_id,omitempty"`

	// RDS数据库信息。
	RdsObjInfo *string `json:"rds_obj_info,omitempty"`

	// DWS数据库信息。
	DwsObjInfo *string `json:"dws_obj_info,omitempty"`

	// 云数据库信息，该字段已废弃。
	ClouddbObjInfo *string `json:"clouddb_obj_info,omitempty"`
}

func (o DataBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBase struct{}"
	}

	return strings.Join([]string{"DataBase", string(data)}, " ")
}
