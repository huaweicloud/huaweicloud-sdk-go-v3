package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityGroupPolicyRequestBody struct {

	// **参数解释** 安全组策略的唯一标识，用于指定待更新的目标安全组策略 **约束限制** 需确保该策略已存在于指定集群和命名空间下，否则返回策略不存在错误 **取值范围** 字符长度1-64位，支持UUID格式（32位十六进制字符，含4个短横线分隔） **默认取值** 无
	PolicyId string `json:"policy_id"`

	// **参数解释** 安全组策略的名称，用于标识策略用途，更新时可修改该名称 **约束限制** 名称不能包含特殊字符（如@、#、$等），且同一命名空间下策略名称建议唯一 **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线（-）、下划线（_） **默认取值** 无，不修改策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释** 待关联到安全组策略的安全组集合，更新后策略将仅应用于该列表中的安全组 **约束限制** 数组不能为空（至少包含1个安全组），且安全组需已存在于当前项目/企业项目下 **取值范围** 数组长度1-20个元素，每个元素需符合SecurityGroup对象结构 **默认取值** 无
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

func (o UpdateSecurityGroupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupPolicyRequestBody", string(data)}, " ")
}
