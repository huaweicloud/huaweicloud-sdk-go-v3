package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisReportForRocketMqRequest Request Object
type ShowDiagnosisReportForRocketMqRequest struct {

	// **参数解释**： 报告ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportId string `json:"report_id"`
}

func (o ShowDiagnosisReportForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisReportForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisReportForRocketMqRequest", string(data)}, " ")
}
