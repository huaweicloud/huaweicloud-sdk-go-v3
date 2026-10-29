package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TreeNode 节点信息
type TreeNode struct {

	// **参数解释:** 目录或文件名。 **取值范围:** 不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释:** 目录或文件路径。 **取值范围:** 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释:** 是否为叶子节点。 **取值范围:** - true：叶子节点。 - false：非叶子节点。
	IsLeaf *bool `json:"is_leaf,omitempty"`

	// **参数解释:** 屏蔽状态。 **取值范围:** - unchecked：不屏蔽。 - all：全屏蔽。 - half: 半屏蔽。
	CheckboxStatus *string `json:"checkbox_status,omitempty"`
}

func (o TreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeNode struct{}"
	}

	return strings.Join([]string{"TreeNode", string(data)}, " ")
}
