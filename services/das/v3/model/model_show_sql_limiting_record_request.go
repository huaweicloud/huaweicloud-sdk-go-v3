package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingRecordRequest Request Object
type ShowSqlLimitingRecordRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// SQL类型
	SqlType *string `json:"sql_type,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 查询ID
	QueryId *string `json:"query_id,omitempty"`

	// 页码
	CurPage *string `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *string `json:"per_page,omitempty"`
}

func (o ShowSqlLimitingRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingRecordRequest", string(data)}, " ")
}
