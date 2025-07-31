package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TotalNum **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
type TotalNum struct {
}

func (o TotalNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalNum struct{}"
	}

	return strings.Join([]string{"TotalNum", string(data)}, " ")
}
