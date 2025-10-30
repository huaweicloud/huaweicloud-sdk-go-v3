package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportSubId **参数解释**: 报告子ID **取值范围**: 字符长度10-2147483647位
type ReportSubId struct {
}

func (o ReportSubId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportSubId struct{}"
	}

	return strings.Join([]string{"ReportSubId", string(data)}, " ")
}
