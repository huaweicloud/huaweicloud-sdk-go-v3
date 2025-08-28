package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangePolicyGroupRequestInfo struct {

	// **参数解释**: 需要修改的的策略组的策略组ID **约束限制**: 需要使用 ListPolicyGroup 接口查询策略组ID，在 ListPolicyGroup 接口的响应体中返回的 group_id 是可以被修改的策略组ID **取值范围**: 字符长度36-64 **默认取值**: 不涉及
	GroupId string `json:"group_id"`

	// **参数解释**: 防护模式 **约束限制**: 不涉及 **取值范围**: - high_detection: 高检出模式。 - equalization: 均衡模式。  **默认取值**: 不涉及
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o ChangePolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangePolicyGroupRequestInfo", string(data)}, " ")
}
