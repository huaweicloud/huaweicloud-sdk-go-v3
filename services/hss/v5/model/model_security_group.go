package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroup struct {

	// **参数解释** 云原生网络模型中安全组的唯一标识，用于关联具体安全组到策略 **约束限制** 安全组需与集群处于同一VPC网络，否则关联失败 **取值范围** 字符长度1-64位，支持字母、数字、短横线（-）、下划线（_） **默认取值** 无
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释** 安全组的名称，用于辅助标识安全组，仅作展示用途 **约束限制** 若传入该参数，需与security_group_id对应的安全组名称一致，否则可能导致展示异常（不影响功能） **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线（-）、下划线（_） **默认取值** 无，默认使用安全组ID对应的系统名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// **参数解释** 安全组的描述信息，用于记录安全组的用途、权限范围等备注 **约束限制** 描述内容不能包含HTML标签等特殊字符 **取值范围** 字符长度0-256位，支持中文、英文、数字、常用标点符号及空格 **默认取值** 无，不修改安全组描述（若原有描述为空则保持为空）
	SecurityGroupDescription *string `json:"security_group_description,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
