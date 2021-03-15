package model

import (
	"encoding/json"

	"strings"
)

// 命令参数响应对象。
type ServiceCommandResponse struct {
	// 设备命令响应名称。
	ResponseName string `json:"response_name"`
	// 设备命令响应的参数列表。
	Paras *[]ServiceCommandPara `json:"paras,omitempty"`
}

func (o ServiceCommandResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceCommandResponse struct{}"
	}

	return strings.Join([]string{"ServiceCommandResponse", string(data)}, " ")
}
