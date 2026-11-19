package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportIdRes **参数解释**: 报告ID **取值范围**: 最小值0，最大值2147483647
type ReportIdRes struct {
}

func (o ReportIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportIdRes struct{}"
	}

	return strings.Join([]string{"ReportIdRes", string(data)}, " ")
}
