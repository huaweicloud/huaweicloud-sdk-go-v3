package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilePath **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
type FilePath struct {
}

func (o FilePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilePath struct{}"
	}

	return strings.Join([]string{"FilePath", string(data)}, " ")
}
