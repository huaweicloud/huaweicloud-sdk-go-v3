package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationLogicTableVoDetailData 返回数据。
type AggregationLogicTableVoDetailData struct {
	Value *AggregationLogicTableVo `json:"value,omitempty"`
}

func (o AggregationLogicTableVoDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationLogicTableVoDetailData struct{}"
	}

	return strings.Join([]string{"AggregationLogicTableVoDetailData", string(data)}, " ")
}
