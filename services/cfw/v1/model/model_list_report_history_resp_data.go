package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportHistoryRespData **参数解释**： 查询报告发送历史数据 **取值范围**： 不涉及
type ListReportHistoryRespData struct {

	// **参数解释**： 数量限制 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 发送历史列表 **取值范围**： 不涉及
	Records *[]ReportHistoryVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o ListReportHistoryRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportHistoryRespData struct{}"
	}

	return strings.Join([]string{"ListReportHistoryRespData", string(data)}, " ")
}
