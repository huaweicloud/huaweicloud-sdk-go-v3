package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessUsername **参数解释**： 运行进程的用户名 **取值范围**： 字符长度1-128位
type ProcessUsername struct {
}

func (o ProcessUsername) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessUsername struct{}"
	}

	return strings.Join([]string{"ProcessUsername", string(data)}, " ")
}
