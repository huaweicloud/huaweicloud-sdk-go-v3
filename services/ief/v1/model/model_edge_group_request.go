package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点组参数
type EdgeGroupRequest struct {

	// 节点组名称。只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name string `json:"name"`

	// 节点组描述。最大长度255个字符
	Description *string `json:"description,omitempty"`

	// 节点组绑定的节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 节点组标签
	Tags *[]Attributes `json:"tags,omitempty"`
}

func (o EdgeGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGroupRequest struct{}"
	}

	return strings.Join([]string{"EdgeGroupRequest", string(data)}, " ")
}
