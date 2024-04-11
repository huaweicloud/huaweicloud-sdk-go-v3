package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAggregationLogicTablesResponse Response Object
type ListAggregationLogicTablesResponse struct {
	Data           *AggregationLogicTableVoSearchResultData `json:"data,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListAggregationLogicTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregationLogicTablesResponse struct{}"
	}

	return strings.Join([]string{"ListAggregationLogicTablesResponse", string(data)}, " ")
}
