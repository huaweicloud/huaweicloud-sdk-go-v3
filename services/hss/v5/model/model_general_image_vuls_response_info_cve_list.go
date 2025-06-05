package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeneralImageVulsResponseInfoCveList struct {

	// CVE ID
	CveId *string `json:"cve_id,omitempty"`

	// CVSS分值
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o GeneralImageVulsResponseInfoCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralImageVulsResponseInfoCveList struct{}"
	}

	return strings.Join([]string{"GeneralImageVulsResponseInfoCveList", string(data)}, " ")
}
