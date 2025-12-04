package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CveAllowlistItem CVE漏洞列表
type CveAllowlistItem struct {

	// CVE的ID, 比如：CVE-2019-10164
	CveId *string `json:"cve_id,omitempty"`
}

func (o CveAllowlistItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CveAllowlistItem struct{}"
	}

	return strings.Join([]string{"CveAllowlistItem", string(data)}, " ")
}
