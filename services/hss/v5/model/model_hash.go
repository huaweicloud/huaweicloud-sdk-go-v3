package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Hash **参数解释**: 事件白名单SHA256 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
type Hash struct {
}

func (o Hash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hash struct{}"
	}

	return strings.Join([]string{"Hash", string(data)}, " ")
}
