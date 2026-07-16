package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodesResultMsg struct {

	// **参数解释**：解锁的资源池节点ID，取值自节点详情metadata.name字段的值。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：资源池节点批量操作结果。 **取值范围**：可选值如下： - failed：操作失败。 - success：操作成功。
	Status *string `json:"status,omitempty"`
}

func (o NodesResultMsg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesResultMsg struct{}"
	}

	return strings.Join([]string{"NodesResultMsg", string(data)}, " ")
}
