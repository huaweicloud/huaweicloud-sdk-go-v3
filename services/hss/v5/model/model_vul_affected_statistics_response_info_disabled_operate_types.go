package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectedStatisticsResponseInfoDisabledOperateTypes **参数解释**: 禁用的漏洞操作类型列表
type VulAffectedStatisticsResponseInfoDisabledOperateTypes struct {

	// **参数解释**: 禁用的操作类型 **取值范围**: - immediate_repair：修复
	OperateType *string `json:"operate_type,omitempty"`

	// **参数解释**: 禁用原因 **取值范围**: 字符长度0-4096位
	Reason *string `json:"reason,omitempty"`
}

func (o VulAffectedStatisticsResponseInfoDisabledOperateTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectedStatisticsResponseInfoDisabledOperateTypes struct{}"
	}

	return strings.Join([]string{"VulAffectedStatisticsResponseInfoDisabledOperateTypes", string(data)}, " ")
}
