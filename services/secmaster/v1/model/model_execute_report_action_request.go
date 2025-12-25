package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteReportActionRequest Request Object
type ExecuteReportActionRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 报告ID
	ReportId string `json:"report_id"`

	Body *ExecuteReportActionInfo `json:"body,omitempty"`
}

func (o ExecuteReportActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteReportActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteReportActionRequest", string(data)}, " ")
}
