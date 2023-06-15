package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 只读节点权重信息
type ModifyProxyWeightReadonlyNode struct {

	// 只读节点ID
	Id *string `json:"id,omitempty"`

	// 只读节点权重
	Weight *int32 `json:"weight,omitempty"`
}

func (o ModifyProxyWeightReadonlyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProxyWeightReadonlyNode struct{}"
	}

	return strings.Join([]string{"ModifyProxyWeightReadonlyNode", string(data)}, " ")
}
