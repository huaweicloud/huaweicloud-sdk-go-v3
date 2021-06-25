package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLogContentsResponse struct {
	// 日志信息。

	Logs           *[]LogContents `json:"logs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateLogContentsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContentsResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogContentsResponse", string(data)}, " ")
}
