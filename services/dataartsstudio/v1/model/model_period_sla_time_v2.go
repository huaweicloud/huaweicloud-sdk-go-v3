package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodSlaTimeV2 struct {

	// 周期序号。
	PeriodNum *int64 `json:"period_num,omitempty"`

	// 相应周期的时间。
	PeriodTime *string `json:"period_time,omitempty"`
}

func (o PeriodSlaTimeV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodSlaTimeV2 struct{}"
	}

	return strings.Join([]string{"PeriodSlaTimeV2", string(data)}, " ")
}
