package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodObject struct {

	// 周期，取值为1到24。
	PeriodNum int32 `json:"period_num"`

	// 小时:分钟，样例：00:59
	PeriodTime string `json:"period_time"`
}

func (o PeriodObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodObject struct{}"
	}

	return strings.Join([]string{"PeriodObject", string(data)}, " ")
}
