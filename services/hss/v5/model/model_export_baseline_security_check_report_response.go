package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBaselineSecurityCheckReportResponse Response Object
type ExportBaselineSecurityCheckReportResponse struct {

	// **参数解释**： 任务ID **取值范围**： 不涉及
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportBaselineSecurityCheckReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBaselineSecurityCheckReportResponse struct{}"
	}

	return strings.Join([]string{"ExportBaselineSecurityCheckReportResponse", string(data)}, " ")
}
