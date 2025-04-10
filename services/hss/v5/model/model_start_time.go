package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTime 启动时间，毫秒
type StartTime struct {
}

func (o StartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTime struct{}"
	}

	return strings.Join([]string{"StartTime", string(data)}, " ")
}
