package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComparisonOperator 阈值符号, 支持的值为(>|<|>=|<=|=|><)
type ComparisonOperator struct {
}

func (o ComparisonOperator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComparisonOperator struct{}"
	}

	return strings.Join([]string{"ComparisonOperator", string(data)}, " ")
}
