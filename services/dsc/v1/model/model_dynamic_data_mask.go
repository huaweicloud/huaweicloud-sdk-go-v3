package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DynamicDataMask struct {

	// 脱敏策略列表，每一个策略对应一个字段，脱敏策略数最多100个。
	MaskStrategies []MaskStrategies `json:"mask_strategies" xml:"mask_strategies"`

	// 数据列表。
	Data []map[string]interface{} `json:"data" xml:"data"`
}

func (o DynamicDataMask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicDataMask struct{}"
	}

	return strings.Join([]string{"DynamicDataMask", string(data)}, " ")
}
