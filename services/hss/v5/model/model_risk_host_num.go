package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHostNum **参数解释**: 有风险服务器数 **取值范围**: 0到2147483647
type RiskHostNum struct {
}

func (o RiskHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHostNum struct{}"
	}

	return strings.Join([]string{"RiskHostNum", string(data)}, " ")
}
