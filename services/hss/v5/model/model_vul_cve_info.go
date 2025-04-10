package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulCveInfo cve 信息
type VulCveInfo struct {

	// cve id
	CveId *string `json:"cve_id,omitempty"`

	// cve评分
	Cvss *float32 `json:"cvss,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`
}

func (o VulCveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulCveInfo struct{}"
	}

	return strings.Join([]string{"VulCveInfo", string(data)}, " ")
}
