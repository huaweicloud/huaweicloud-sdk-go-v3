package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendQueryDataResp struct {

	// 查询时间。
	QueryTime *int64 `json:"query_time,omitempty"`

	// 监控指标名称。
	IndicatorName *string `json:"indicator_name,omitempty"`

	// 监控对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 次级监控id。
	SubObjectId *string `json:"sub_object_id,omitempty"`

	// 节点数据。
	DataPoints *[]TrendQueryData `json:"data_points,omitempty"`
}

func (o TrendQueryDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendQueryDataResp struct{}"
	}

	return strings.Join([]string{"TrendQueryDataResp", string(data)}, " ")
}
