package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppPathRes **参数解释**: 软件路径 **取值范围**: 字符长度0-512位
type AppPathRes struct {
}

func (o AppPathRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppPathRes struct{}"
	}

	return strings.Join([]string{"AppPathRes", string(data)}, " ")
}
