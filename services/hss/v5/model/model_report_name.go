package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportName **参数解释**: 报告名称 **取值范围**: 字符长度1-128位
type ReportName struct {
}

func (o ReportName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportName struct{}"
	}

	return strings.Join([]string{"ReportName", string(data)}, " ")
}
