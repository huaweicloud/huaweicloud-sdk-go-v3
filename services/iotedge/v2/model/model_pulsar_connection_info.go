package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarConnectionInfo 创建外部推送通道请求结构体
type PulsarConnectionInfo struct {

	// 鉴权token
	Token string `json:"token"`
}

func (o PulsarConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarConnectionInfo struct{}"
	}

	return strings.Join([]string{"PulsarConnectionInfo", string(data)}, " ")
}
