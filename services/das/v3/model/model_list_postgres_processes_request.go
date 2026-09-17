package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPostgresProcessesRequest Request Object
type ListPostgresProcessesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

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

	// 是否显示全部
	ShowAll *bool `json:"show_all,omitempty"`

	// 是否显示没有后台进程的会话
	ShowNoPid *bool `json:"show_no_pid,omitempty"`

	// 指定慢sql阈值
	Time *string `json:"time,omitempty"`

	// 页码
	CurPage *string `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *string `json:"per_page,omitempty"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序方式（asc/desc）
	Order *string `json:"order,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点类型
	NodeRole *string `json:"node_role,omitempty"`

	// 是否过滤系统会话
	HideSys *bool `json:"hide_sys,omitempty"`
}

func (o ListPostgresProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresProcessesRequest struct{}"
	}

	return strings.Join([]string{"ListPostgresProcessesRequest", string(data)}, " ")
}
