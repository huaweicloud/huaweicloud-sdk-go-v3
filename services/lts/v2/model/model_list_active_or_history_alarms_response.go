package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListActiveOrHistoryAlarmsResponse struct {

	// 告警信息
	Events *[]Events `json:"events,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListActiveOrHistoryAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsResponse", string(data)}, " ")
}
