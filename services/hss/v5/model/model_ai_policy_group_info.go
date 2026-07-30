package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiPolicyGroupInfo struct {

	// **参数解释**： 策略组ID **取值范围**： 字符长度1-20位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 策略组名称 **取值范围**: 字符长度1-128位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 最小值0，最大值2147483647
	GroupType *int32 `json:"group_type,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 防护对象 **取值范围**： 字符长度1-128位
	ProtectionObject *string `json:"protection_object,omitempty"`

	// **参数解释**: 防护对象类型 **取值范围**: - 0：云服务 - 1：三方
	ObjectType *int32 `json:"object_type,omitempty"`

	// **参数解释**: 防护对象个数 **取值范围**: 取值0-100000
	ObjectNum *int32 `json:"object_num,omitempty"`

	// **参数解释**: 是否是默认策略 **取值范围**: - false：否 - true：是
	IsDefault *bool `json:"is_default,omitempty"`

	// **参数解释**: 是否是默认策略 **取值范围**: - false：否 - true：是
	IsExclusive *bool `json:"is_exclusive,omitempty"`

	// **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
	DetailIsUsed *bool `json:"detail_is_used,omitempty"`

	// **参数解释**: 描述 **取值范围**: 字符长度0-256位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 策略列表 **取值范围**: 不涉及
	PolicyList *[]AiPolicyList `json:"policy_list,omitempty"`

	// **参数解释**: 智能体列表 **取值范围**: 不涉及
	AgentIdList *[]string `json:"agent_id_list,omitempty"`
}

func (o AiPolicyGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupInfo struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupInfo", string(data)}, " ")
}
