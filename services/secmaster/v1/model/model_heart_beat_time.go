package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HeartBeatTime 最后一次接收到心跳信号的时间
type HeartBeatTime struct {
}

func (o HeartBeatTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeartBeatTime struct{}"
	}

	return strings.Join([]string{"HeartBeatTime", string(data)}, " ")
}
