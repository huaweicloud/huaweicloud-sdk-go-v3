package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resize struct {

	// 调整大小目标规格
	NodeType string `json:"node_type"`

	// 调整大小目标节点数
	NumberOfNode int32 `json:"number_of_node"`
}

func (o Resize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resize struct{}"
	}

	return strings.Join([]string{"Resize", string(data)}, " ")
}
