package model

import (
	"encoding/json"

	"strings"
)

// 容灾演练虚拟机数据结构
type DrillServerParams struct {
	// 演练云服务器对应的保护实例ID。

	ProtectedInstance string `json:"protected_instance"`
	// 演练云服务器ID。

	DrillServerId string `json:"drill_server_id"`
}

func (o DrillServerParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DrillServerParams struct{}"
	}

	return strings.Join([]string{"DrillServerParams", string(data)}, " ")
}
