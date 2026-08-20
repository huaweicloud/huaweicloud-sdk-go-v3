package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUrlSummary top url 详情数据
type TopUrlSummary struct {

	// **参数解释：** URL名称 **取值范围：** 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释：** 对应查询类型的值 **取值范围：** 若值为流量，流量单位：Byte
	Value *int64 `json:"value,omitempty"`

	// **参数解释：** 查询起始时间戳 **取值范围：** 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 查询结束时间戳 **取值范围：** 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 统计指标类型 **取值范围：** - flux：流量 - req_num：请求总数
	StatType *string `json:"stat_type,omitempty"`
}

func (o TopUrlSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrlSummary struct{}"
	}

	return strings.Join([]string{"TopUrlSummary", string(data)}, " ")
}
