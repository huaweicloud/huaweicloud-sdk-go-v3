package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点参数
type EdgeNodeUpdate struct {

	// 边缘节点描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 边缘节点日志配置
	LogConfigs *[]LogConfigs `json:"log_configs,omitempty"`

	// NTP服务器地址，每个节点最多仅能配置两个。D310表示D310类型；D910表示D910类型；不填表示为D310类型。
	NtpServers *[]string `json:"ntp_servers,omitempty"`

	// 边缘节点属性，关联属性个数最多为32个
	Attributes *[]Attributes `json:"attributes,omitempty"`
}

func (o EdgeNodeUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeUpdate struct{}"
	}

	return strings.Join([]string{"EdgeNodeUpdate", string(data)}, " ")
}
