package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CmdLine **参数解释**： 进程命令行 **约束限制**： 不涉及
type CmdLine struct {
}

func (o CmdLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CmdLine struct{}"
	}

	return strings.Join([]string{"CmdLine", string(data)}, " ")
}
