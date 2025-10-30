package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportId **参数解释**: 报告ID **取值范围**: 字符长度10-2147483647位
type ReportId struct {
}

func (o ReportId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportId struct{}"
	}

	return strings.Join([]string{"ReportId", string(data)}, " ")
}
