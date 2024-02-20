package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalyzerId 分析器的唯一标识符。
type AnalyzerId struct {
}

func (o AnalyzerId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerId struct{}"
	}

	return strings.Join([]string{"AnalyzerId", string(data)}, " ")
}
