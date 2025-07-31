package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnprotectHostNum **参数解释**: 未防护服务器数 **取值范围**: 0到2147483647
type UnprotectHostNum struct {
}

func (o UnprotectHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnprotectHostNum struct{}"
	}

	return strings.Join([]string{"UnprotectHostNum", string(data)}, " ")
}
