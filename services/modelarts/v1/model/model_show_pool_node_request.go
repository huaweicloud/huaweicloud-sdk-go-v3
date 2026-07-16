package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolNodeRequest Request Object
type ShowPoolNodeRequest struct {

	// **参数解释**：资源池的ID，取值自资源池详情的metadata.name字段。 **约束限制**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为36-63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：节点名称。取值节点详情的metadata.name字段。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeName string `json:"node_name"`
}

func (o ShowPoolNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolNodeRequest", string(data)}, " ")
}
