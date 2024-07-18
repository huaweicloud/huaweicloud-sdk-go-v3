package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecAppTaskResultResponse Response Object
type ShowSecAppTaskResultResponse struct {
	BasicInfo *BasicInfo `json:"basic_info,omitempty"`

	ApkComponentInfo *ApkComponentInfo `json:"apk_component_info,omitempty"`

	HapComponentInfo *HapComponentInfo `json:"hap_component_info,omitempty"`

	// 漏洞列表
	VulnList *[]VulnInfo `json:"vuln_list,omitempty"`

	// 隐私合规列表
	PrivacyComplianceList *[]PrivacyComplianceInfo `json:"privacy_compliance_list,omitempty"`
	HttpStatusCode        int                      `json:"-"`
}

func (o ShowSecAppTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecAppTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSecAppTaskResultResponse", string(data)}, " ")
}
