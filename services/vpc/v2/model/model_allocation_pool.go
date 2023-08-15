package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllocationPool
type AllocationPool struct {

	// 网络池结束IP
	End *string `json:"end,omitempty"`

	// 网络池起始IP
	Start *string `json:"start,omitempty"`
}

func (o AllocationPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllocationPool struct{}"
	}

	return strings.Join([]string{"AllocationPool", string(data)}, " ")
}
