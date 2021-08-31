package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCaseRequest struct {
	// 用例id

	CaseId int32 `json:"case_id"`
}

func (o DeleteCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteCaseRequest", string(data)}, " ")
}
