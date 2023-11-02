package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentDiagnosisInstanceInfosPerMetricResponse Response Object
type ShowIntelligentDiagnosisInstanceInfosPerMetricResponse struct {

	// 当前指标的异常实例总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 异常实例信息。
	InstanceInfos  *[]InstanceInfoForDiagnosis `json:"instance_infos,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowIntelligentDiagnosisInstanceInfosPerMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentDiagnosisInstanceInfosPerMetricResponse struct{}"
	}

	return strings.Join([]string{"ShowIntelligentDiagnosisInstanceInfosPerMetricResponse", string(data)}, " ")
}
