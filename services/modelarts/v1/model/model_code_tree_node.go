package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTreeNode 算法目录树节点。
type CodeTreeNode struct {

	// **参数解释**：算法目录树当前层级目录名。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：算法目录树当前层级目录下子文件和子目录。
	Children *[]CodeTreeNode `json:"children,omitempty"`
}

func (o CodeTreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTreeNode struct{}"
	}

	return strings.Join([]string{"CodeTreeNode", string(data)}, " ")
}
