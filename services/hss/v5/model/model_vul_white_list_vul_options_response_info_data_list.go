package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulWhiteListVulOptionsResponseInfoDataList struct {

	// **参数解释**: 漏洞id **取值范围**: 字符长度0-256
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256
	VulName *string `json:"vul_name,omitempty"`
}

func (o VulWhiteListVulOptionsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulWhiteListVulOptionsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulWhiteListVulOptionsResponseInfoDataList", string(data)}, " ")
}
