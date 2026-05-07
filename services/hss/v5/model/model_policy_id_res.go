package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyIdRes **参数解释**: 策略ID **取值范围**: 字符长度1-64位
type PolicyIdRes struct {
}

func (o PolicyIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyIdRes struct{}"
	}

	return strings.Join([]string{"PolicyIdRes", string(data)}, " ")
}
