package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbsConnectionRequestBody DBS连接请求体
type CreateDbsConnectionRequestBody struct {

	// 数据库引擎类型，取值范围：mysql, sqlserver, postgresql, taurus, gaussdbv5, mongodb, ddm
	EngineType string `json:"engine_type"`

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 数据库来源类型，取值范围：rds, gaussdb, dds, ddm
	NetworkType string `json:"network_type"`

	// 用户名
	Username string `json:"username"`

	// 是否保存密码
	IsSavePassword bool `json:"is_save_password"`

	// 密码
	Password string `json:"password"`

	// 节点ID列表，实例节点的唯一标识
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 端口，取值范围：[1,65536]
	Port *int32 `json:"port,omitempty"`

	// 数据库名字
	DatabaseName *string `json:"database_name,omitempty"`

	// SQL记录开关
	SqlRecordFlag *bool `json:"sql_record_flag,omitempty"`
}

func (o CreateDbsConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbsConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDbsConnectionRequestBody", string(data)}, " ")
}
