package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipeAlias 管道别名
type PipeAlias struct {
}

func (o PipeAlias) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeAlias struct{}"
	}

	return strings.Join([]string{"PipeAlias", string(data)}, " ")
}
