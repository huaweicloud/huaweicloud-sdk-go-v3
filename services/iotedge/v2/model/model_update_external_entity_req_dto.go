package model

import (
	"encoding/json"

	"strings"
)

// 更新外部实体请求结构体
type UpdateExternalEntityReqDto struct {
	// 连接外部实体的协议类型

	Protocol string `json:"protocol"`
	// 连接类型

	ConnectionType string `json:"connection_type"`

	MqttConnectionInfo *MqttConnectionInfo `json:"mqtt_connection_info,omitempty"`
}

func (o UpdateExternalEntityReqDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateExternalEntityReqDto struct{}"
	}

	return strings.Join([]string{"UpdateExternalEntityReqDto", string(data)}, " ")
}
