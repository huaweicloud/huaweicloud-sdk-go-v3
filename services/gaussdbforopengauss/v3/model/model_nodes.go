package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点信息
type Nodes struct {

	// 节点ID。
	Id string `json:"id"`

	// 组件列表。
	Components []Components `json:"components"`
}

func (o Nodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nodes struct{}"
	}

	return strings.Join([]string{"Nodes", string(data)}, " ")
}
