package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnknownNum **参数解释**: 识别未知进程数 **取值范围**: 最小值0，最大值2147483647
type UnknownNum struct {
}

func (o UnknownNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnknownNum struct{}"
	}

	return strings.Join([]string{"UnknownNum", string(data)}, " ")
}
