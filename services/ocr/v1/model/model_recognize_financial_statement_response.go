package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeFinancialStatementResponse Response Object
type RecognizeFinancialStatementResponse struct {
	Result         *FinancialStatementResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RecognizeFinancialStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFinancialStatementResponse struct{}"
	}

	return strings.Join([]string{"RecognizeFinancialStatementResponse", string(data)}, " ")
}
