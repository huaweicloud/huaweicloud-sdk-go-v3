package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccumulatedTimes 累计次数
type AccumulatedTimes struct {
}

func (o AccumulatedTimes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccumulatedTimes struct{}"
	}

	return strings.Join([]string{"AccumulatedTimes", string(data)}, " ")
}
