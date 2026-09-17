package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRequestBody 创建实例连接请求体
type CreateConnectionRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库来源类型
	NetworkType string `json:"network_type"`

	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`

	// 是否保存密码
	IsSavePassword bool `json:"is_save_password"`

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// SQL记录开关
	SqlRecordFlag *bool `json:"sql_record_flag,omitempty"`
}

func (o CreateConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequestBody", string(data)}, " ")
}
