package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyConnectionNewRequestBody 验证数据库实例连接请求体
type VerifyConnectionNewRequestBody struct {

	// 数据库引擎类型，取值范围：mysql, sqlserver, postgresql, taurus, gaussdbv5, mongodb, ddm
	EngineType string `json:"engine_type"`

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 数据库来源类型，取值范围：rds, gaussdb, dds, ddm
	NetworkType string `json:"network_type"`

	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`

	// 节点ID列表，实例节点的唯一标识
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 端口，取值范围：[1,65536]
	Port *int32 `json:"port,omitempty"`

	// 数据库名字
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o VerifyConnectionNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyConnectionNewRequestBody struct{}"
	}

	return strings.Join([]string{"VerifyConnectionNewRequestBody", string(data)}, " ")
}
