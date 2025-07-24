package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapacityBytes 容量（单位：byte）
type CapacityBytes struct {
}

func (o CapacityBytes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityBytes struct{}"
	}

	return strings.Join([]string{"CapacityBytes", string(data)}, " ")
}
