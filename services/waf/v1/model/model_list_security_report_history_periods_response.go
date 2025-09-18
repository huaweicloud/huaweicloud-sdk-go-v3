package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportHistoryPeriodsResponse Response Object
type ListSecurityReportHistoryPeriodsResponse struct {

	// **参数解释：** 总数，安全报告历史统计周期记录的总条数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 条目，安全报告历史统计周期的详细列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]ListSecurityReportHistoryPeriodResponseItems `json:"items,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ListSecurityReportHistoryPeriodsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportHistoryPeriodsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityReportHistoryPeriodsResponse", string(data)}, " ")
}
