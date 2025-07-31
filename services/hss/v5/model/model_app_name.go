package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppName **参数解释**: 软件名称 **取值范围**: 字符长度1-256位
type AppName struct {
}

func (o AppName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppName struct{}"
	}

	return strings.Join([]string{"AppName", string(data)}, " ")
}
