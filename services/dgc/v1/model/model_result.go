package model

import (
	"encoding/json"

	"strings"
)

type Result struct {
	Message  *string `json:"message,omitempty"`
	RowCount *int32  `json:"rowCount,omitempty"`
	Rows     *string `json:"rows,omitempty"`
	Schema   *string `json:"schema,omitempty"`
}

func (o Result) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
