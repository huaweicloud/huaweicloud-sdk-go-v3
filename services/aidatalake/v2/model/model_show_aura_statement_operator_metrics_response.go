package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraStatementOperatorMetricsResponse Response Object
type ShowAuraStatementOperatorMetricsResponse struct {
	Body           *[]AuraStatementOperatorMetrics `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowAuraStatementOperatorMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraStatementOperatorMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowAuraStatementOperatorMetricsResponse", string(data)}, " ")
}
