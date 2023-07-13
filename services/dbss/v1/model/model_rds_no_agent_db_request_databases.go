package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsNoAgentDbRequestDatabases struct {

	// 数据库ID
	Id string `json:"id"`

	// 数据库名称
	DbName string `json:"db_name"`

	// 数据库状态
	Status string `json:"status"`

	// 数据库端口
	Port string `json:"port"`

	// 数据库IP
	Ip string `json:"ip"`

	// 数据库实例名称
	InstanceName string `json:"instance_name"`

	// 数据库版本
	Version string `json:"version"`

	// 数据库类型
	Type string `json:"type"`

	// 企业项目ID
	EnterpriseId string `json:"enterprise_id"`

	// 企业项目名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
}

func (o RdsNoAgentDbRequestDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsNoAgentDbRequestDatabases struct{}"
	}

	return strings.Join([]string{"RdsNoAgentDbRequestDatabases", string(data)}, " ")
}
