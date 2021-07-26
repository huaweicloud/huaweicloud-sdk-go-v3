package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLtsLogDetailsResponse struct {
	// 日志组id

	GroupId *string `json:"group_id,omitempty"`
	// 日志流id

	StreamId *string `json:"stream_id,omitempty"`
	// 日志流名称

	StreamName     *string `json:"stream_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLtsLogDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLtsLogDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowLtsLogDetailsResponse", string(data)}, " ")
}
