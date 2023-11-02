package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse Response Object
type ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse struct {

	// 诊断信息列表。
	DiagnosisInfo  *[]DiagnosisInfo `json:"diagnosis_info,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse", string(data)}, " ")
}
