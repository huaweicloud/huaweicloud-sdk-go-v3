package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitItemResp **参数解释** 单位 **取值范围** 长度为[0,32]个字符
type UnitItemResp struct {
}

func (o UnitItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitItemResp struct{}"
	}

	return strings.Join([]string{"UnitItemResp", string(data)}, " ")
}
