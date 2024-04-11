package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignAggregationLogicTableResponse Response Object
type CreateDesignAggregationLogicTableResponse struct {
	Data           *AggregationLogicTableVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o CreateDesignAggregationLogicTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignAggregationLogicTableResponse struct{}"
	}

	return strings.Join([]string{"CreateDesignAggregationLogicTableResponse", string(data)}, " ")
}
