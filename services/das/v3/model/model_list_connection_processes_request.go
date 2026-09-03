package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionProcessesRequest Request Object
type ListConnectionProcessesRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 指定用户
	User *string `json:"user,omitempty"`

	// 指定访问来源
	Host *string `json:"host,omitempty"`

	// 指定数据库
	Db *string `json:"db,omitempty"`

	// 指定状态
	State *string `json:"state,omitempty"`

	// 指定命令
	Command *string `json:"command,omitempty"`

	// 模糊搜索条件
	Keywords *string `json:"keywords,omitempty"`

	// 是否显示全部
	ShowAll *bool `json:"show_all,omitempty"`

	// 是否显示没有后台进程的会话
	ShowNoPid *bool `json:"show_no_pid,omitempty"`

	// 指定慢sql阈值
	Time *string `json:"time,omitempty"`

	// 每页记录数
	PerPage *string `json:"per_page,omitempty"`

	// 页码
	CurPage *string `json:"cur_page,omitempty"`

	// 选择排序列
	OrderBy *string `json:"order_by,omitempty"`

	// 排序顺序
	Order *string `json:"order,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点类型
	NodeRole *string `json:"node_role,omitempty"`

	// 是否过滤系统会话
	HideSys *bool `json:"hide_sys,omitempty"`
}

func (o ListConnectionProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionProcessesRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionProcessesRequest", string(data)}, " ")
}
