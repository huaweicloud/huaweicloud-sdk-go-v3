package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 限制条目数
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}
