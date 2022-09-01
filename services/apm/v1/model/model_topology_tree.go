package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopologyTree struct {

	// 拓扑树节点
	Tree *[]TreeNode `json:"tree,omitempty" xml:"tree"`
}

func (o TopologyTree) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyTree struct{}"
	}

	return strings.Join([]string{"TopologyTree", string(data)}, " ")
}
