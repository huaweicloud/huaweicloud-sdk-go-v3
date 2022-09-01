package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodesWeight struct {

	// 数据库节点id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 权重。取值范围：0~1000。
	Weight *int32 `json:"weight,omitempty" xml:"weight"`
}

func (o NodesWeight) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesWeight struct{}"
	}

	return strings.Join([]string{"NodesWeight", string(data)}, " ")
}
