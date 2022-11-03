package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivacyComplianceInfo struct {

	// 隐私合规类型
	Category *string `json:"category,omitempty"`

	// 隐私合规子类型列表
	SubtypeList *[]PrivacyComplianceSubtype `json:"subtype_list,omitempty"`
}

func (o PrivacyComplianceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivacyComplianceInfo struct{}"
	}

	return strings.Join([]string{"PrivacyComplianceInfo", string(data)}, " ")
}
