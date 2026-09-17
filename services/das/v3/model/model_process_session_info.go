package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessSessionInfo 会话信息
type ProcessSessionInfo struct {

	// 会话ID
	Id *int64 `json:"id,omitempty"`

	// 数据库用户
	User *string `json:"user,omitempty"`

	// 数据库主机
	Host *string `json:"host,omitempty"`

	// 数据库名称
	Db *string `json:"db,omitempty"`

	// 命令类型
	Command *string `json:"command,omitempty"`

	// SQL信息
	SqlInfo *string `json:"sql_info,omitempty"`

	// 当前状态持续时间
	StateDuration *int64 `json:"state_duration,omitempty"`

	// 当前状态
	State *string `json:"state,omitempty"`
}

func (o ProcessSessionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessSessionInfo struct{}"
	}

	return strings.Join([]string{"ProcessSessionInfo", string(data)}, " ")
}
