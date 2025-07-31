package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrustNum **参数解释**: 识别可信进程数 **取值范围**: 最小值0，最大值2147483647
type TrustNum struct {
}

func (o TrustNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustNum struct{}"
	}

	return strings.Join([]string{"TrustNum", string(data)}, " ")
}
