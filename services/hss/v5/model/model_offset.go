package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Offset 偏移量：指定返回记录的开始位置
type Offset struct {
}

func (o Offset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Offset struct{}"
	}

	return strings.Join([]string{"Offset", string(data)}, " ")
}
