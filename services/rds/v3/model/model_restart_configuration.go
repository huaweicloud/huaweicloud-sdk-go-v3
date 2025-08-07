package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartConfiguration struct {

	// 是否重启虚拟机。
	RestartServer *bool `json:"restart_server,omitempty"`

	// 是否强制重启, 强制重启会导致数据库服务中未提交的事务强制中断。
	Forcible *bool `json:"forcible,omitempty"`

	// 是否在可维护时间段内重启。
	Delay *bool `json:"delay,omitempty"`
}

func (o RestartConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartConfiguration struct{}"
	}

	return strings.Join([]string{"RestartConfiguration", string(data)}, " ")
}
