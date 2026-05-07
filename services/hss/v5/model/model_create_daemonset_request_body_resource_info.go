package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDaemonsetRequestBodyResourceInfo 资源配置信息
type CreateDaemonsetRequestBodyResourceInfo struct {

	// **参数解释**： 资源限制类型：默认规则or自定义or自适应 **约束限制**： 不涉及 **取值范围**： - default：默认类型。 - customized：用户自定义类型。 - adaptive：自适应类型。  **默认取值**： 不涉及
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: cpu最大值 **约束限制**: 不涉及 **取值范围**: 字符长度0-32位 **默认取值**: 不涉及
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存最大值 **约束限制**: 不涉及 **取值范围**: 字符长度0-32位 **默认取值**: 不涉及
	MemLimit *string `json:"mem_limit,omitempty"`
}

func (o CreateDaemonsetRequestBodyResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDaemonsetRequestBodyResourceInfo struct{}"
	}

	return strings.Join([]string{"CreateDaemonsetRequestBodyResourceInfo", string(data)}, " ")
}
