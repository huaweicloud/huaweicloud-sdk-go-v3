package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignAggregationLogicTableResponse Response Object
type UpdateDesignAggregationLogicTableResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateDesignAggregationLogicTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignAggregationLogicTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignAggregationLogicTableResponse", string(data)}, " ")
}
