package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RaspCveInfo CVE信息
type RaspCveInfo struct {

	// **参数解释**: CVE编号 **取值范围**: 字符长度0-64位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: CVSS评分 **取值范围**: 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o RaspCveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RaspCveInfo struct{}"
	}

	return strings.Join([]string{"RaspCveInfo", string(data)}, " ")
}
