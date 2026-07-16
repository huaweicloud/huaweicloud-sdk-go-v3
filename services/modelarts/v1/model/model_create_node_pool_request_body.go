package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodePoolRequestBody 创建资源池请求体。
type CreateNodePoolRequestBody struct {

	// **参数解释**：API版本。 **取值范围**：可选值如下： - v2
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：节点池类型。 **取值范围**：可选值如下： -  NodePool：节点池
	Kind string `json:"kind"`

	Metadata *CreateNodePoolMetaVo `json:"metadata"`

	Spec *NodePoolSpec `json:"spec,omitempty"`
}

func (o CreateNodePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNodePoolRequestBody", string(data)}, " ")
}
