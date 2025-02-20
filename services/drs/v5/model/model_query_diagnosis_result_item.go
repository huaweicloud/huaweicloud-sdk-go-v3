package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDiagnosisResultItem 诊断项。
type QueryDiagnosisResultItem struct {

	// 诊断项类别。
	Category *string `json:"category,omitempty"`

	// 诊断项名称。
	Name *string `json:"name,omitempty"`

	// 诊断项权重。
	Weight *float64 `json:"weight,omitempty"`
}

func (o QueryDiagnosisResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnosisResultItem struct{}"
	}

	return strings.Join([]string{"QueryDiagnosisResultItem", string(data)}, " ")
}
