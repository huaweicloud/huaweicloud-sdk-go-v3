package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessFilename **参数解释**： 进程文件名 **取值范围**： 字符长度1-64位
type ProcessFilename struct {
}

func (o ProcessFilename) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessFilename struct{}"
	}

	return strings.Join([]string{"ProcessFilename", string(data)}, " ")
}
