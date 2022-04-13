package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 通道信息
type ChannelCreateInfo struct {
	// 通道名称，字符串长度4-24，必须包含a-z，0-9，以小写字母开头，以小写字母或者数字结尾

	ChannelName string `json:"channel_name"`
	// 通道描述

	ChannelDescription *string `json:"channel_description,omitempty"`
}

func (o ChannelCreateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelCreateInfo struct{}"
	}

	return strings.Join([]string{"ChannelCreateInfo", string(data)}, " ")
}
