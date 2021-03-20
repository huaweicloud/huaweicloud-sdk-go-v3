package model

import (
	"encoding/json"

	"strings"
)

// 命令服务对象。
type ServiceCommand struct {
	// 设备命令名称。

	CommandName string `json:"command_name"`
	// 设备命令的参数列表。

	Paras *[]ServiceCommandPara `json:"paras,omitempty"`
	// 设备命令的响应列表。

	Responses *[]ServiceCommandResponse `json:"responses,omitempty"`
}

func (o ServiceCommand) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceCommand struct{}"
	}

	return strings.Join([]string{"ServiceCommand", string(data)}, " ")
}
