package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserHomeDir **参数解释**： 用户目录 **取值范围**： 字符长度1-256位
type UserHomeDir struct {
}

func (o UserHomeDir) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserHomeDir struct{}"
	}

	return strings.Join([]string{"UserHomeDir", string(data)}, " ")
}
