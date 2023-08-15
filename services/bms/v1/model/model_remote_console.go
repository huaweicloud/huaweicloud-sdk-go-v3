package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteConsole
type RemoteConsole struct {

	// 远程登录的协议。
	Protocol string `json:"protocol"`

	// 远程登录的类型。
	Type string `json:"type"`
}

func (o RemoteConsole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteConsole struct{}"
	}

	return strings.Join([]string{"RemoteConsole", string(data)}, " ")
}
