package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilePath 文件路径
type FilePath struct {
}

func (o FilePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilePath struct{}"
	}

	return strings.Join([]string{"FilePath", string(data)}, " ")
}
