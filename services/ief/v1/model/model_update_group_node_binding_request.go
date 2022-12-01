package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 绑定或解绑节点请求体
type UpdateGroupNodeBindingRequest struct {

	// 新增绑定的节点ID列表
	AddNodeIds *[]string `json:"add_node_ids,omitempty"`

	// 新增解绑的节点ID列表
	RemoveNodeIds *[]string `json:"remove_node_ids,omitempty"`
}

func (o UpdateGroupNodeBindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupNodeBindingRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupNodeBindingRequest", string(data)}, " ")
}
