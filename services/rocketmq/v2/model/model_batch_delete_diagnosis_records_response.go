package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDiagnosisRecordsResponse Response Object
type BatchDeleteDiagnosisRecordsResponse struct {

	// **参数解释**： 诊断报告ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportIdList   *[]string `json:"report_id_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteDiagnosisRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDiagnosisRecordsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDiagnosisRecordsResponse", string(data)}, " ")
}
