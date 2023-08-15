package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHostNum 有风险服务器数
type RiskHostNum struct {
}

func (o RiskHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHostNum struct{}"
	}

	return strings.Join([]string{"RiskHostNum", string(data)}, " ")
}
