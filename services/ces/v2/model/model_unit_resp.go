package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitResp **参数解释**： 数据的单位。    **取值范围**： 长度为[0,32]个字符。
type UnitResp struct {
}

func (o UnitResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitResp struct{}"
	}

	return strings.Join([]string{"UnitResp", string(data)}, " ")
}
