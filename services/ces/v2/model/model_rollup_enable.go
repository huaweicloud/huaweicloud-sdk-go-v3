package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupEnable 是否开启聚合
type RollupEnable struct {
}

func (o RollupEnable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupEnable struct{}"
	}

	return strings.Join([]string{"RollupEnable", string(data)}, " ")
}
