package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulUrgentResponseInfoDisabledOperateTypes struct {

	// **参数解释**: 操作类型 **取值范围**: - ignore: 忽略 - not_ignore: 取消忽略 - immediate_repair: 修复 - manual_repair: 人工修复 - verify: 验证 - add_to_whitelist: 加入白名单
	OperateType *string `json:"operate_type,omitempty"`

	// **参数解释**: 不可进行操作的原因 **取值范围**: 字符范围0-512位
	Reason *string `json:"reason,omitempty"`
}

func (o VulUrgentResponseInfoDisabledOperateTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulUrgentResponseInfoDisabledOperateTypes struct{}"
	}

	return strings.Join([]string{"VulUrgentResponseInfoDisabledOperateTypes", string(data)}, " ")
}
