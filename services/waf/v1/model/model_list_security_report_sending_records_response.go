package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportSendingRecordsResponse Response Object
type ListSecurityReportSendingRecordsResponse struct {

	// **参数解释：** 总数，安全报告发送记录的总条数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 条目，安全报告发送记录的详细列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]ListSecurityReportSendingRecordResponseItems `json:"items,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ListSecurityReportSendingRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportSendingRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityReportSendingRecordsResponse", string(data)}, " ")
}
