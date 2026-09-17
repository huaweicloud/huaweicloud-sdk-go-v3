package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceProcessesRequest Request Object
type ListInstanceProcessesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 用户名
	User *string `json:"user,omitempty"`

	// 访问来源IP
	Host *string `json:"host,omitempty"`

	// 数据库
	Db *string `json:"db,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 命令
	Command *string `json:"command,omitempty"`

	// 模糊搜索条件
	Keywords *string `json:"keywords,omitempty"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序方式（asc/desc）
	Order *string `json:"order,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 数据库来源类型
	NetworkType *string `json:"network_type,omitempty"`
}

func (o ListInstanceProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceProcessesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceProcessesRequest", string(data)}, " ")
}
