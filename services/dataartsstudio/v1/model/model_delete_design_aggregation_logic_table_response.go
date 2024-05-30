package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignAggregationLogicTableResponse Response Object
type DeleteDesignAggregationLogicTableResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignAggregationLogicTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignAggregationLogicTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignAggregationLogicTableResponse", string(data)}, " ")
}
