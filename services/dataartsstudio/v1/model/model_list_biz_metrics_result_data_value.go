package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricsResultDataValue value，统一的返回结果的外层数据结构。
type ListBizMetricsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// BizMetricVO信息。
	Records *[]BizMetricVo `json:"records,omitempty"`
}

func (o ListBizMetricsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListBizMetricsResultDataValue", string(data)}, " ")
}
