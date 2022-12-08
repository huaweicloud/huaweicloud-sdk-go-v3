package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息
type DataBase struct {

	// 数据库ID
	Id *string `json:"id,omitempty"`

	// 数据库名称
	Name string `json:"name"`

	// 添加的数据库类型： 枚举值：  MYSQL  ORACLE  POSTGRESQL  SQLSERVER  DAMENG  TAURUS  DWS  KINGBASE  GAUSSDBOPENGAUSS   GREENPLUM   HIGHGO   SHENTONG   GBASE8A   GBASE8S   GBASEXDM   MONGODB   DDS
	Type string `json:"type"`

	// 数据库版本
	Version string `json:"version"`

	// 数据库字符集
	Charset string `json:"charset"`

	// 数据库IP
	Ip string `json:"ip"`

	// 数据库端口
	Port string `json:"port"`

	// 数据库操作系统
	Os string `json:"os"`

	// 开启状态（1：开启，0：关闭）
	Status *string `json:"status,omitempty"`

	// 数据库实例名
	InstanceName string `json:"instance_name"`

	// 数据库的运行状态 枚举值：  ACTIVE  SHUTOFF  ERROR
	AuditStatus *string `json:"audit_status,omitempty"`

	// agent的唯一ID
	AgentUrl *[]string `json:"agent_url,omitempty"`

	// 数据库分类，取值范围： RDS（表示RDS数据库）和 ECS（自建数据库）
	DbClassification string `json:"db_classification"`
}

func (o DataBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBase struct{}"
	}

	return strings.Join([]string{"DataBase", string(data)}, " ")
}
