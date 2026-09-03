package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseObjectsRequest Request Object
type ListDatabaseObjectsRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// Schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 表ID
	TableId *string `json:"table_id,omitempty"`

	// 开始时间(Unix timestamp),单位:毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间(Unix timestamp),单位:毫秒
	EndAt int64 `json:"end_at"`

	// 页码
	PageNum int32 `json:"page_num"`

	// 每页记录数
	PageSize int32 `json:"page_size"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序方式（asc/desc）
	Order *string `json:"order,omitempty"`

	// 额外排序字段
	ExtraOrderBy *string `json:"extra_order_by,omitempty"`

	// 额外排序方式
	ExtraOrder *string `json:"extra_order,omitempty"`

	// 对象类型
	ObjType *string `json:"obj_type,omitempty"`

	// 返回类型
	RetType *string `json:"ret_type,omitempty"`

	// 是否系统对象
	IsSys *string `json:"is_sys,omitempty"`

	// 对象子类型
	ObjSubType *string `json:"obj_sub_type,omitempty"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	ObjName *string `json:"obj_name,omitempty"`

	Keywords *string `json:"keywords,omitempty"`

	CurPage *string `json:"cur_page,omitempty"`

	PerPage *string `json:"per_page,omitempty"`
}

func (o ListDatabaseObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseObjectsRequest", string(data)}, " ")
}
