package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmDetailRequest struct {

	// 时区
	TimeZone string `json:"time_zone"`

	// 当前页
	Offset *string `json:"offset,omitempty"`

	// 总页数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmDetailRequest", string(data)}, " ")
}
