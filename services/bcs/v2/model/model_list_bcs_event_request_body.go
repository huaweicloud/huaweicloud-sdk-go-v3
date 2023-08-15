package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBcsEventRequestBody 查询和统计事件告警请求体
type ListBcsEventRequestBody struct {

	// timeRange用于指标查询时间范围，主要用于解决客户端时间和服务端时间不一致情况下，查询最近N分钟的数据。另可用于精确查询某一段时间的数据。
	TimeRange string `json:"time_range"`

	// 统计步长。毫秒数
	Step int64 `json:"step"`

	// 模糊查询匹配字段，可以为空
	Search *string `json:"search,omitempty"`

	// 查询条件组合，可以为空
	MetadataRelation *[]EventMetadataRelation `json:"metadata_relation,omitempty"`

	Sort *EventResultSort `json:"sort,omitempty"`
}

func (o ListBcsEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsEventRequestBody struct{}"
	}

	return strings.Join([]string{"ListBcsEventRequestBody", string(data)}, " ")
}
