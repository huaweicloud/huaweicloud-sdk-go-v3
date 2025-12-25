package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChannelReadWriteCount 数值
type ChannelReadWriteCount struct {
}

func (o ChannelReadWriteCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelReadWriteCount struct{}"
	}

	return strings.Join([]string{"ChannelReadWriteCount", string(data)}, " ")
}
