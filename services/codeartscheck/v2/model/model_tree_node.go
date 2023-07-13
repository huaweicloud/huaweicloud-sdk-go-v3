package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TreeNode 节点信息
type TreeNode struct {

	// 目录或文件名
	FileName *string `json:"file_name,omitempty"`

	// 目录或文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 是否为叶子节点，true是，false不是
	IsLeaf *bool `json:"is_leaf,omitempty"`

	// 屏蔽状态，包括unchecked(不屏蔽)、all(全屏蔽)、half(半屏蔽)
	CheckboxStatus *string `json:"checkbox_status,omitempty"`
}

func (o TreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeNode struct{}"
	}

	return strings.Join([]string{"TreeNode", string(data)}, " ")
}
