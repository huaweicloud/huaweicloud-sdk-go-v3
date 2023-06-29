package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemEventsResponse Response Object
type ListSystemEventsResponse struct {

	// 数目
	Count *int32 `json:"count,omitempty"`

	// 系统订阅详情列表
	Systemevents   *[]Event `json:"systemevents,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListSystemEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSystemEventsResponse", string(data)}, " ")
}
