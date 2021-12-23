package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulnsLevel struct {
	// 高危漏洞数

	High *int32 `json:"high,omitempty"`
	// 中危漏洞数

	Middle *int32 `json:"middle,omitempty"`
	// 低危漏洞数

	Low *int32 `json:"low,omitempty"`
	// 提示危漏洞数

	Hint *int32 `json:"hint,omitempty"`
}

func (o VulnsLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnsLevel struct{}"
	}

	return strings.Join([]string{"VulnsLevel", string(data)}, " ")
}
