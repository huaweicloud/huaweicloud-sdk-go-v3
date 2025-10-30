package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulWhiteListDetailResponseInfoCves struct {

	// **参数解释**: cve id **取值范围**: 字符长度0-256
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: cve的cvss评分 **取值范围**: 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o VulWhiteListDetailResponseInfoCves) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulWhiteListDetailResponseInfoCves struct{}"
	}

	return strings.Join([]string{"VulWhiteListDetailResponseInfoCves", string(data)}, " ")
}
