package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAopWorkflowInfo 更新流程的请求信息
type ModifyAopWorkflowInfo struct {

	// **参数解释**: 流程名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程的描述 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 启用或者禁用流程 **约束限制**: 不涉及         **取值范围**: - true  启用流程 - false 禁用流程  **默认值**:  不涉及
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 流程版本ID **约束限制**: 当前流程下已经激活的流程版本ID，当启用流程时为必填参数
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**: 流程实体类型 - IP IP - ACCOUNT 账号 - DOMAIN 域名  **约束限制**: 不涉及         **取值范围**: - IP - ACCOUNT - DOMAIN  **默认值**:  不涉及
	Labels *string `json:"labels,omitempty"`
}

func (o ModifyAopWorkflowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAopWorkflowInfo struct{}"
	}

	return strings.Join([]string{"ModifyAopWorkflowInfo", string(data)}, " ")
}
