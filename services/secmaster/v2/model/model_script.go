package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Script 脚本内容，长度在1到10240个字符之间。
type Script struct {
}

func (o Script) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Script struct{}"
	}

	return strings.Join([]string{"Script", string(data)}, " ")
}
