package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptDescription 脚本的相关描述信息，长度在1到1024个字符之间。
type ScriptDescription struct {
}

func (o ScriptDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptDescription struct{}"
	}

	return strings.Join([]string{"ScriptDescription", string(data)}, " ")
}
