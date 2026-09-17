package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePulsarNodeChannelConnectionInfo 更新外部推送通道请求结构体
type UpdatePulsarNodeChannelConnectionInfo struct {

	// 鉴权token
	Token *string `json:"token,omitempty"`
}

func (o UpdatePulsarNodeChannelConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePulsarNodeChannelConnectionInfo struct{}"
	}

	return strings.Join([]string{"UpdatePulsarNodeChannelConnectionInfo", string(data)}, " ")
}
