package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationLogicTableVoSearchResultData 返回数据。
type AggregationLogicTableVoSearchResultData struct {
	Value *AggregationLogicTableVoSearchResultDataValue `json:"value,omitempty"`
}

func (o AggregationLogicTableVoSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationLogicTableVoSearchResultData struct{}"
	}

	return strings.Join([]string{"AggregationLogicTableVoSearchResultData", string(data)}, " ")
}
