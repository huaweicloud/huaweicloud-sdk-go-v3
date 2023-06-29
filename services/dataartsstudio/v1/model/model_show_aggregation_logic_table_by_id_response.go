package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregationLogicTableByIdResponse Response Object
type ShowAggregationLogicTableByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAggregationLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregationLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregationLogicTableByIdResponse", string(data)}, " ")
}
