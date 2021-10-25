package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeFinancialStatementResponse struct {
	Result         *FinancialStatementResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RecognizeFinancialStatementResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeFinancialStatementResponse struct{}"
	}

	return strings.Join([]string{"RecognizeFinancialStatementResponse", string(data)}, " ")
}
