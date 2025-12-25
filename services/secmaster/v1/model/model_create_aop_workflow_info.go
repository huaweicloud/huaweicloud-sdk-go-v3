package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAopWorkflowInfo 创建流程的请求信息
type CreateAopWorkflowInfo struct {

	// **参数解释**: 流程名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name string `json:"name"`

	// **参数解释**: 流程的描述 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 引擎的类型 - public_engine 共享版  **约束限制**: 不涉及         **取值范围**: - public_engine  **默认值**:  不涉及
	EngineType string `json:"engine_type"`

	// **参数解释**: 流程的类型 - NORMAL 通用 - SURVEY 调查 - HEMOSTASIS 止血 - EASE 缓解  **约束限制**: 不涉及         **取值范围**: - NORMAL - SURVEY - HEMOSTASIS - EASE  **默认值**:  不涉及
	AopType string `json:"aop_type"`

	// **参数解释**: 数据类的ID **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	DataclassId string `json:"dataclass_id"`

	// **参数解释**: 流程实体类型 - IP IP - ACCOUNT 账号 - DOMAIN 域名  **约束限制**: 不涉及         **取值范围**: - IP - ACCOUNT - DOMAIN  **默认值**:  不涉及
	Labels *string `json:"labels,omitempty"`
}

func (o CreateAopWorkflowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAopWorkflowInfo struct{}"
	}

	return strings.Join([]string{"CreateAopWorkflowInfo", string(data)}, " ")
}
