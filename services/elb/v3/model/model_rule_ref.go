package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RuleRef struct {
	// 规则ID。

	Id string `json:"id"`
}

func (o RuleRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRef struct{}"
	}

	return strings.Join([]string{"RuleRef", string(data)}, " ")
}
