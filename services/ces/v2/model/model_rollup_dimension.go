package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupDimension 聚合维度
type RollupDimension struct {
}

func (o RollupDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupDimension struct{}"
	}

	return strings.Join([]string{"RollupDimension", string(data)}, " ")
}
