package model

import (
	"encoding/json"

	"strings"
)

// 命令服务对象。
type ServiceCommand struct {
	// **参数说明**：设备命令名称。 **取值范围**：长度不超过64，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合。

	CommandName string `json:"command_name"`
	// **参数说明**：设备命令的参数列表。

	Paras *[]ServiceCommandPara `json:"paras,omitempty"`
	// **参数说明**：设备命令的响应列表。

	Responses *[]ServiceCommandResponse `json:"responses,omitempty"`
}

func (o ServiceCommand) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceCommand struct{}"
	}

	return strings.Join([]string{"ServiceCommand", string(data)}, " ")
}
