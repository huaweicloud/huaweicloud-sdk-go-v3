package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEventResponse struct {
	// 攻击事件数量

	Total *int32 `json:"total,omitempty"`
	// 攻击事件详情

	Items          *[]ListEventResponseBodyItems `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListEventResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventResponse struct{}"
	}

	return strings.Join([]string{"ListEventResponse", string(data)}, " ")
}
