package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新终端节点的详细信息。
type UpdateEndpointOption struct {

	// 终端节点权重。
	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateEndpointOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointOption struct{}"
	}

	return strings.Join([]string{"UpdateEndpointOption", string(data)}, " ")
}
