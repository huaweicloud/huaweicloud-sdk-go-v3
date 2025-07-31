package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessPath **参数解释**： 进程路径 **取值范围**： 字符长度1-256位
type ProcessPath struct {
}

func (o ProcessPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessPath struct{}"
	}

	return strings.Join([]string{"ProcessPath", string(data)}, " ")
}
