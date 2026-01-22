package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnosisTaskForRocketMqResponse Response Object
type CreateDiagnosisTaskForRocketMqResponse struct {

	// **参数解释**： 报告ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportId       *string `json:"report_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnosisTaskForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskForRocketMqResponse", string(data)}, " ")
}
