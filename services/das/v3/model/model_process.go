package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Process struct {

	// 会话ID
	Id string `json:"id" xml:"id"`

	// 用户
	User string `json:"user" xml:"user"`

	// 主机
	Host string `json:"host" xml:"host"`

	// 数据库
	Database string `json:"database" xml:"database"`

	// 命令
	Command string `json:"command" xml:"command"`

	// 会话持续时间
	Time string `json:"time" xml:"time"`

	// 状态
	State string `json:"state" xml:"state"`

	// SQL语句
	Sql string `json:"sql" xml:"sql"`

	// 事务持续时间
	TrxExecutedTime string `json:"trx_executed_time" xml:"trx_executed_time"`
}

func (o Process) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Process struct{}"
	}

	return strings.Join([]string{"Process", string(data)}, " ")
}
