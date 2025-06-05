package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVulReportDataResponseInfoCveList struct {

	// CVE ID
	CveId *string `json:"cve_id,omitempty"`

	// CVSS分值
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o ShowVulReportDataResponseInfoCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoCveList struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoCveList", string(data)}, " ")
}
