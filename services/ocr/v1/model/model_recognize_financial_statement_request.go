package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeFinancialStatementRequest struct {
	Body *FinancialStatementRequestBody `json:"body,omitempty"`
}

func (o RecognizeFinancialStatementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFinancialStatementRequest struct{}"
	}

	return strings.Join([]string{"RecognizeFinancialStatementRequest", string(data)}, " ")
}
