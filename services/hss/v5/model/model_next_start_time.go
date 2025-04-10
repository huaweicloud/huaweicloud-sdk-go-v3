package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NextStartTime 下次启动时间，毫秒
type NextStartTime struct {
}

func (o NextStartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextStartTime struct{}"
	}

	return strings.Join([]string{"NextStartTime", string(data)}, " ")
}
