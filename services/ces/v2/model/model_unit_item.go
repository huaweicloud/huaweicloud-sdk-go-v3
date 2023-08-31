package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitItem 单位
type UnitItem struct {
}

func (o UnitItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitItem struct{}"
	}

	return strings.Join([]string{"UnitItem", string(data)}, " ")
}
