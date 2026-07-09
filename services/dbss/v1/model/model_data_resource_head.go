package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataResourceHead struct {

	// 数据库实例别名
	Alias *string `json:"alias,omitempty"`

	Datastore *DbDataStore `json:"datastore,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据库端口
	DbPort *string `json:"db_port,omitempty"`

	// 数据库用户列表
	DbUserList *[]DbUser `json:"db_user_list,omitempty"`

	// 云服务名称，云上数据库服务，如：rds
	Provider *string `json:"provider,omitempty"`

	// rds数据库ID
	RdsId *string `json:"rds_id,omitempty"`

	// 数据库状态
	Status *string `json:"status,omitempty"`

	// 数据库类型
	Type *string `json:"type,omitempty"`
}

func (o DataResourceHead) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataResourceHead struct{}"
	}

	return strings.Join([]string{"DataResourceHead", string(data)}, " ")
}
