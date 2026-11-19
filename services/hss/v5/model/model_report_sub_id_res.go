package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportSubIdRes **参数解释**: 报告子ID **取值范围**: 最小值0，最大值2147483647
type ReportSubIdRes struct {
}

func (o ReportSubIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportSubIdRes struct{}"
	}

	return strings.Join([]string{"ReportSubIdRes", string(data)}, " ")
}
