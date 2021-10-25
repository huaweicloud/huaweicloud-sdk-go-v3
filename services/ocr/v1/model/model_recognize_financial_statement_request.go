package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeFinancialStatementRequest struct {
	Body *FinancialStatementRequestBody `json:"body,omitempty"`
}

func (o RecognizeFinancialStatementRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeFinancialStatementRequest struct{}"
	}

	return strings.Join([]string{"RecognizeFinancialStatementRequest", string(data)}, " ")
}
