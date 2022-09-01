package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点信息
type Nodes struct {

	// 节点ID。
	Id string `json:"id" xml:"id"`

	// 组件列表。
	Components []Components `json:"components" xml:"components"`
}

func (o Nodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nodes struct{}"
	}

	return strings.Join([]string{"Nodes", string(data)}, " ")
}
