package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEventsResponse struct {
	// 一条或者多条事件数据。
	Events         *[]EventInfo   `json:"events,omitempty"`
	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
