package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipeName 管道名称
type PipeName struct {
}

func (o PipeName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeName struct{}"
	}

	return strings.Join([]string{"PipeName", string(data)}, " ")
}
