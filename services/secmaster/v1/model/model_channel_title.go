package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChannelTitle 相关标题
type ChannelTitle struct {
}

func (o ChannelTitle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelTitle struct{}"
	}

	return strings.Join([]string{"ChannelTitle", string(data)}, " ")
}
