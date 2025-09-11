package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupDimensionResp **参数解释** 聚合维度 **取值范围** 长度为[1,32]个字符
type RollupDimensionResp struct {
}

func (o RollupDimensionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupDimensionResp struct{}"
	}

	return strings.Join([]string{"RollupDimensionResp", string(data)}, " ")
}
