package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraStatementQueryMetricsResponse Response Object
type ShowAuraStatementQueryMetricsResponse struct {
	Body           *[]StatementQueryMetrics `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAuraStatementQueryMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraStatementQueryMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowAuraStatementQueryMetricsResponse", string(data)}, " ")
}
