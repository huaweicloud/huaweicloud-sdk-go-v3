package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 只读节点权重信息
type TaurusModifyProxyWeightReadonlyNode struct {

	// 只读节点id
	Id *string `json:"id,omitempty"`

	// 只读节点权重
	Weight *int32 `json:"weight,omitempty"`
}

func (o TaurusModifyProxyWeightReadonlyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusModifyProxyWeightReadonlyNode struct{}"
	}

	return strings.Join([]string{"TaurusModifyProxyWeightReadonlyNode", string(data)}, " ")
}
