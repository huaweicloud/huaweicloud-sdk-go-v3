package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTime 屏蔽起始时间，HH:mm:ss。
type StartTime struct {
}

func (o StartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTime struct{}"
	}

	return strings.Join([]string{"StartTime", string(data)}, " ")
}
