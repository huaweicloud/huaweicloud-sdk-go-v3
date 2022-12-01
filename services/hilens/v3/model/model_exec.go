package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Exec struct {

	// 探针执行命令，最大长度10240个字符。不允许^#~^$|%&*<>()'\"[]{}这些特殊字符
	Command string `json:"command"`
}

func (o Exec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Exec struct{}"
	}

	return strings.Join([]string{"Exec", string(data)}, " ")
}
