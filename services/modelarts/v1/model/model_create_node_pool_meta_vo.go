package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodePoolMetaVo 创建节点池的metadata信息。
type CreateNodePoolMetaVo struct {

	// **参数解释**：节点池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	Annotations *CreateNodePoolAnnotations `json:"annotations,omitempty"`
}

func (o CreateNodePoolMetaVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolMetaVo struct{}"
	}

	return strings.Join([]string{"CreateNodePoolMetaVo", string(data)}, " ")
}
