package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentEventsResponse Response Object
type ListComponentEventsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentEventKindObj `json:"kind,omitempty"`

	// 事件项。
	Items          *[]EventItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListComponentEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentEventsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentEventsResponse", string(data)}, " ")
}
