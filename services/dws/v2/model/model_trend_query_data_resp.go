package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendQueryDataResp struct {

	// **参数解释**： 查询时间。 **取值范围**： 不涉及。
	QueryTime *int64 `json:"query_time,omitempty"`

	// **参数解释**： 监控指标名称。 **取值范围**： 不涉及。
	IndicatorName *string `json:"indicator_name,omitempty"`

	// **参数解释**： 监控对象ID。 **取值范围**： 不涉及。
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： 单位。 **取值范围**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 次级监控ID。 **取值范围**： 不涉及。
	SubObjectId *string `json:"sub_object_id,omitempty"`

	// **参数解释**： 节点数据。 **取值范围**： 不涉及。
	DataPoints *[]TrendQueryData `json:"data_points,omitempty"`
}

func (o TrendQueryDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendQueryDataResp struct{}"
	}

	return strings.Join([]string{"TrendQueryDataResp", string(data)}, " ")
}
