package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisReportRequest Request Object
type ShowDiagnosisReportRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 报告ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportId string `json:"report_id"`
}

func (o ShowDiagnosisReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisReportRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisReportRequest", string(data)}, " ")
}
