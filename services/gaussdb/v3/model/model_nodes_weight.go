package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodesWeight struct {

	// 数据库节点ID。
	Id *string `json:"id,omitempty"`

	// 权重。取值范围：0~1000。
	Weight *int32 `json:"weight,omitempty"`
}

func (o NodesWeight) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesWeight struct{}"
	}

	return strings.Join([]string{"NodesWeight", string(data)}, " ")
}
