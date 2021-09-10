package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExecuteTbSessionResponse struct {
	// 所有数据的信息。

	Questions      *[]ExecuteTbQuestion `json:"questions,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ExecuteTbSessionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteTbSessionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTbSessionResponse", string(data)}, " ")
}
