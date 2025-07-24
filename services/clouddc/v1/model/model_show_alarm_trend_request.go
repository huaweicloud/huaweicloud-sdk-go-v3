package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmTrendRequest Request Object
type ShowAlarmTrendRequest struct {

	// 查询开始时间戳，单位秒，from 必须小于 to
	From string `json:"from"`

	// 查询结束时间戳，单位秒，
	To string `json:"to"`
}

func (o ShowAlarmTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmTrendRequest", string(data)}, " ")
}
