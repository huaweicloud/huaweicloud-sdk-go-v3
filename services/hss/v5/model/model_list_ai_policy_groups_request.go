package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPolicyGroupsRequest Request Object
type ListAiPolicyGroupsRequest struct {

	// **参数解释**: 策略组ID **约束限制**: 不涉及 **取值范围**: 字符长度1-20位 **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 策略组名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 防护对象 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	ProtectionObject *string `json:"protection_object,omitempty"`

	// **参数解释**: 对象类型 **约束限制**: 不涉及 **取值范围**: - 0：云服务 - 1：三方  **默认取值**: 不涉及
	ObjectType *int32 `json:"object_type,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAiPolicyGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPolicyGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListAiPolicyGroupsRequest", string(data)}, " ")
}
