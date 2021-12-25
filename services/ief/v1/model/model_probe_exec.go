package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行命令检查，与http_get二选一
type ProbeExec struct {
	// 探针执行命令，最大长度10240个字符

	Command string `json:"command"`
}

func (o ProbeExec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProbeExec struct{}"
	}

	return strings.Join([]string{"ProbeExec", string(data)}, " ")
}
