package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeSpans struct {

	// 开始时间；不包含日期，为本地时刻，根据网关侧的本地时间进行调度
	Start string `json:"start"`

	// 结束时间，不包含日期，为本地时刻，根据网关侧的本地时间进行调度
	End string `json:"end"`
}

func (o TimeSpans) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeSpans struct{}"
	}

	return strings.Join([]string{"TimeSpans", string(data)}, " ")
}
