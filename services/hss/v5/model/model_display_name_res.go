package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisplayNameRes **参数解释**: 主题的显示名 **取值范围**: 字符长度0-256位
type DisplayNameRes struct {
}

func (o DisplayNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayNameRes struct{}"
	}

	return strings.Join([]string{"DisplayNameRes", string(data)}, " ")
}
