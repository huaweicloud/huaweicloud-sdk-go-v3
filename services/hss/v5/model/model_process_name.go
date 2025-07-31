package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessName **参数解释**： 进程名称 **取值范围**： 字符长度1-128位
type ProcessName struct {
}

func (o ProcessName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessName struct{}"
	}

	return strings.Join([]string{"ProcessName", string(data)}, " ")
}
