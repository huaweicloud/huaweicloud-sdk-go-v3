package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRaspEventsResponse Response Object
type ListRaspEventsResponse struct {

	// total number
	TotalNum *int64 `json:"total_num,omitempty"`

	// data list
	DataList       *[]RaspProtectHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListRaspEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRaspEventsResponse", string(data)}, " ")
}
