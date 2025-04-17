package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddFullSqlTaskBody struct {

	// 数据库引擎
	DatastoreType string `json:"datastore_type"`

	// 节点ID，有值时创建节点维度任务
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt int64 `json:"end_at"`

	// 用户名列表，最大长度20
	UserList *[]string `json:"user_list,omitempty"`

	// 关键字列表，最大长度20
	KeywordList *[]string `json:"keyword_list,omitempty"`

	// 数据库列表，最大长度20
	DbList *[]string `json:"db_list,omitempty"`

	// 操作类型列表，最大长度20
	OperationList *[]string `json:"operation_list,omitempty"`

	// 线程id列表，最大长度20
	ThreadIdList *[]string `json:"thread_id_list,omitempty"`

	// 执行状态列表，\"0\"为成功，\"1\"为失败，最大长度20
	StatusList *[]string `json:"status_list,omitempty"`
}

func (o AddFullSqlTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFullSqlTaskBody struct{}"
	}

	return strings.Join([]string{"AddFullSqlTaskBody", string(data)}, " ")
}
