package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaurusDbProcessInfo TaurusDB用户会话线程信息
type TaurusDbProcessInfo struct {

	// **参数解释**：  用户会话线程ID。
	Id int64 `json:"id"`

	// **参数解释**：  启动用户会话线程的用户。
	User string `json:"user"`

	// **参数解释**：  发送请求的主机和端口。
	Host string `json:"host"`

	// **参数解释**：  当前访问的数据库名。
	Db string `json:"db"`

	// **参数解释**：  当前执行的命令。
	Command string `json:"command"`

	// **参数解释**：  用户会话线程处于当前状态的持续时间，单位为秒。
	Time int64 `json:"time"`

	// **参数解释**：  正在执行的SQL语句的当前状态。
	State string `json:"state"`

	// **参数解释**：  额外信息，通常是正在执行的语句。
	Info string `json:"info"`
}

func (o TaurusDbProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusDbProcessInfo struct{}"
	}

	return strings.Join([]string{"TaurusDbProcessInfo", string(data)}, " ")
}
