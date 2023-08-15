package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentEventsResponse Response Object
type ListComponentEventsResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“ComponentEvent”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
