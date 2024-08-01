package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PhysicalProcessInfo struct {

	// 会话id
	Id *int64 `json:"id,omitempty"`

	// 用户
	User *string `json:"user,omitempty"`

	// 主机
	Host *string `json:"host,omitempty"`

	// 数据库
	Db *string `json:"db,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 连接状态，一般是休眠（sleep），查询（query），连接（connect）
	Command *string `json:"command,omitempty"`

	// 线程执行的 sql 语句
	Info *string `json:"info,omitempty"`

	// 会话持续时间，单位秒
	Time *int64 `json:"time,omitempty"`

	// 事务持续时间，单位秒
	TrxExecutedTime *int64 `json:"trx_executed_time,omitempty"`
}

func (o PhysicalProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalProcessInfo struct{}"
	}

	return strings.Join([]string{"PhysicalProcessInfo", string(data)}, " ")
}
