package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleIdReq **参数解释**： 规则ID **约束限制**： 必填 **取值范围**： 字符长度1-36位 **默认取值**： 不涉及
type RuleIdReq struct {
}

func (o RuleIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleIdReq struct{}"
	}

	return strings.Join([]string{"RuleIdReq", string(data)}, " ")
}
