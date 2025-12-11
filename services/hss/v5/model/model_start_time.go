package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTime **参数解释**： 启动时间 **取值范围**： 最小值0，最大值9223372036854775807；时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算）；单位：ms
type StartTime struct {
}

func (o StartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTime struct{}"
	}

	return strings.Join([]string{"StartTime", string(data)}, " ")
}
