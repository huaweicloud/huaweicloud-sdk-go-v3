package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupEnable **参数解释** 是否开启聚合 **约束限制** 当RollupEnable开启时，RollupFilter和RollupDimension必填 **取值范围** true，表示开启聚合；false表示不开启聚合 **默认取值** false
type RollupEnable struct {
}

func (o RollupEnable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupEnable struct{}"
	}

	return strings.Join([]string{"RollupEnable", string(data)}, " ")
}
