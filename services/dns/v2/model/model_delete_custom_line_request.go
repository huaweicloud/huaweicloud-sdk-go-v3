package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCustomLineRequest struct {
	// 解析线路ID。

	LineId string `json:"line_id"`
}

func (o DeleteCustomLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCustomLineRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomLineRequest", string(data)}, " ")
}
