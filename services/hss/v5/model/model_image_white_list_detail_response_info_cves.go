package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageWhiteListDetailResponseInfoCves struct {

	// **参数解释**: CVE ID **取值范围**: 字符长度0-256位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: CVE评分 **取值范围**: 取值0-10
	Cvss *float32 `json:"cvss,omitempty"`
}

func (o ImageWhiteListDetailResponseInfoCves) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWhiteListDetailResponseInfoCves struct{}"
	}

	return strings.Join([]string{"ImageWhiteListDetailResponseInfoCves", string(data)}, " ")
}
