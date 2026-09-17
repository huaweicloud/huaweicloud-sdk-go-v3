package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PgProcessInfo PostgreSQL进程信息
type PgProcessInfo struct {

	// 会话ID
	Id *string `json:"id,omitempty"`

	// 用户
	User *string `json:"user,omitempty"`

	// 连接库的IP和port
	Host *string `json:"host,omitempty"`

	// 数据库
	Db *string `json:"db,omitempty"`

	// 当前执行的命令
	Command *string `json:"command,omitempty"`

	// 会话运行时间
	Time *string `json:"time,omitempty"`

	// 执行状态
	State *string `json:"state,omitempty"`

	// 执行的SQL
	Info *string `json:"info,omitempty"`

	// 事务持续时间
	TrxDuration *string `json:"trx_duration,omitempty"`
}

func (o PgProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PgProcessInfo struct{}"
	}

	return strings.Join([]string{"PgProcessInfo", string(data)}, " ")
}
