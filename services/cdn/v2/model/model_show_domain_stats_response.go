package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStatsResponse Response Object
type ShowDomainStatsResponse struct {

	// **参数解释：** 数据分组方式 **取值范围：** domain：按域名分组 **默认取值：** 不分组
	GroupBy *string `json:"group_by,omitempty"`

	// **参数解释：** 查询起始时间戳 **取值范围：** 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 查询结束时间戳 **取值范围：** 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 统计指标类型 **取值范围：** - flux：流量 - req_num：请求总数
	StatType *string `json:"stat_type,omitempty"`

	// **参数解释：** 查询数据类型 **取值范围：** - summary：汇总数据 - detail：明细数据
	Action *string `json:"action,omitempty"`

	// **参数解释：** 查询时间粒度 **取值范围：** - 300：采样时间间隔为5分钟，单位：秒 - 3600：采样时间间隔为1小时，单位：秒 - 86400：采样时间间隔为1天，单位：秒 **默认取值：** 默认取对应时间跨度的最小间隔 > 时间跨度小于等于7天，最小时间间隔为300；时间跨度大于7天，最小时间间隔为3600
	Interval *int64 `json:"interval,omitempty"`

	// **参数解释：** 按指定的分组方式组织的数据 **取值范围：** 不涉及
	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsResponse", string(data)}, " ")
}
