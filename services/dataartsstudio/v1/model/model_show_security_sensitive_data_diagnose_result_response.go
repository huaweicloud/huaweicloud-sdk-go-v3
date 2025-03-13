package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecuritySensitiveDataDiagnoseResultResponse Response Object
type ShowSecuritySensitiveDataDiagnoseResultResponse struct {

	// 诊断任务id
	TaskId *string `json:"task_id,omitempty"`

	// 是否正在扫描
	Scanning *bool `json:"scanning,omitempty"`

	// 最新检测时间。
	CheckTime *int64 `json:"check_time,omitempty"`

	Classification *ClassificationResult `json:"classification,omitempty"`

	Rule *IdentificationRuleResult `json:"rule,omitempty"`

	Masking        *DataMaskingResult `json:"masking,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSecuritySensitiveDataDiagnoseResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecuritySensitiveDataDiagnoseResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSecuritySensitiveDataDiagnoseResultResponse", string(data)}, " ")
}
