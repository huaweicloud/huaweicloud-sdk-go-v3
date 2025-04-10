package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IgnoreDir 排除目录，多个用;分隔
type IgnoreDir struct {
}

func (o IgnoreDir) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnoreDir struct{}"
	}

	return strings.Join([]string{"IgnoreDir", string(data)}, " ")
}
