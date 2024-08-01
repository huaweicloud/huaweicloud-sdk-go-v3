package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogicalProcessInfo struct {

	// 会话id
	Id *string `json:"id,omitempty"`

	// 当前连接用户
	User *string `json:"user,omitempty"`

	// 所属的 IP 和端口
	Host *string `json:"host,omitempty"`

	// 数据库名
	Db *string `json:"db,omitempty"`

	// 连接状态，一般是休眠（sleep），查询（query），连接（connect）
	Command *string `json:"command,omitempty"`

	// 连接状态持续的时间，单位是秒（s）
	Time *string `json:"time,omitempty"`

	// 当前 SQL 语句的状态
	State *string `json:"state,omitempty"`

	// 当前所执行的 SQL 语句
	Info *string `json:"info,omitempty"`
}

func (o LogicalProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalProcessInfo struct{}"
	}

	return strings.Join([]string{"LogicalProcessInfo", string(data)}, " ")
}
