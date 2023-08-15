package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OptionalFlavorsInfo struct {

	// 实例规格变更时可用的规格列表。
	List []ComputeFlavor `json:"list"`

	// 总记录数。
	TotalCount int32 `json:"total_count"`
}

func (o OptionalFlavorsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionalFlavorsInfo struct{}"
	}

	return strings.Join([]string{"OptionalFlavorsInfo", string(data)}, " ")
}
