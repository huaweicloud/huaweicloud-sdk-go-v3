package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Process struct {
	// 会话ID

	Id string `json:"id"`
	// 用户

	User string `json:"user"`
	// 主机

	Host string `json:"host"`
	// 数据库

	Database string `json:"database"`
	// 命令

	Command string `json:"command"`
	// 会话持续时间

	Time string `json:"time"`
	// 状态

	State string `json:"state"`
	// SQL语句

	Sql string `json:"sql"`
	// 事务持续时间

	TrxExecutedTime string `json:"trx_executed_time"`
}

func (o Process) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Process struct{}"
	}

	return strings.Join([]string{"Process", string(data)}, " ")
}
