package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulAffectedStatisticsResponseInfoCceDisabledVulList struct {

	// **参数解释**: 主机名称 **取值范围**: 字符长度0-64位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机id **取值范围**: 字符长度0-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞补丁编号 **取值范围**: 字符长度0-256
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 操作提示 **取值范围**: 字符长度0-4096
	OperationDescription *string `json:"operation_description,omitempty"`
}

func (o VulAffectedStatisticsResponseInfoCceDisabledVulList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectedStatisticsResponseInfoCceDisabledVulList struct{}"
	}

	return strings.Join([]string{"VulAffectedStatisticsResponseInfoCceDisabledVulList", string(data)}, " ")
}
