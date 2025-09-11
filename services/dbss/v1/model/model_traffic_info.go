package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrafficInfo struct {

	// 记录ID
	Id *string `json:"id,omitempty"`

	// 每秒钟接收字节数
	Rxmegabytes *float64 `json:"rxmegabytes,omitempty"`

	// 记录时间
	Time *string `json:"time,omitempty"`

	// 每秒钟发送字节数
	Txmegabytes *float64 `json:"txmegabytes,omitempty"`
}

func (o TrafficInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficInfo struct{}"
	}

	return strings.Join([]string{"TrafficInfo", string(data)}, " ")
}
