package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OptionalFlavorsInfo struct {

	// 扩容节点时可用的规格列表。
	OptionalFlavorList *[]ComputeFlavor `json:"optional_flavor_list,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o OptionalFlavorsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionalFlavorsInfo struct{}"
	}

	return strings.Join([]string{"OptionalFlavorsInfo", string(data)}, " ")
}
