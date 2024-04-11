package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregationLogicTableByIdResponse Response Object
type ShowAggregationLogicTableByIdResponse struct {
	Data           *AggregationLogicTableVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowAggregationLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregationLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregationLogicTableByIdResponse", string(data)}, " ")
}
