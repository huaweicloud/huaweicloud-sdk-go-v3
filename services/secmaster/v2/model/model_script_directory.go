package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptDirectory 脚本目录分组名称，长度在1到256个字符之间。
type ScriptDirectory struct {
}

func (o ScriptDirectory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptDirectory struct{}"
	}

	return strings.Join([]string{"ScriptDirectory", string(data)}, " ")
}
