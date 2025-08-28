package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVulReportDataResponseInfoCveList struct {

	// **参数解释**： CVE ID **取值范围**： 字符长度1-32位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**： CVSS分值 **取值范围**： 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o ShowVulReportDataResponseInfoCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoCveList struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoCveList", string(data)}, " ")
}
