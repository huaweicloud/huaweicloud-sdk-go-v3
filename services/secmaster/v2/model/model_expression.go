package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Expression 表达式
type Expression struct {
}

func (o Expression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Expression struct{}"
	}

	return strings.Join([]string{"Expression", string(data)}, " ")
}
