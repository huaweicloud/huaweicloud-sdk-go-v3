package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHistoryRunInfoResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// log_list

	LogList        *[]HistoryRunInfo `json:"log_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowHistoryRunInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHistoryRunInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryRunInfoResponse", string(data)}, " ")
}
