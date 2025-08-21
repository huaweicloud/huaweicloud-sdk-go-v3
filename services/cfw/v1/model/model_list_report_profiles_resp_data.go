package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportProfilesRespData **参数解释**： 查询模板列表响应 **取值范围**： 不涉及
type ListReportProfilesRespData struct {

	// **参数解释**： 数量限制 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 模板列表 **取值范围**： 不涉及
	Records *[]ReportProfileVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o ListReportProfilesRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportProfilesRespData struct{}"
	}

	return strings.Join([]string{"ListReportProfilesRespData", string(data)}, " ")
}
