package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretEventsResponse Response Object
type ListSecretEventsResponse struct {

	// 事件详情列表。
	Events *[]Event `json:"events,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecretEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretEventsResponse", string(data)}, " ")
}
