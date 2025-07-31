package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostNum **参数解释**: 影响主机数量 **取值范围**: 最小值0，最大值2147483647
type HostNum struct {
}

func (o HostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNum struct{}"
	}

	return strings.Join([]string{"HostNum", string(data)}, " ")
}
