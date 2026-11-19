package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportCategoryRes **参数解释**: 报告类别 **取值范围**: - daily_report：安全日报 - weekly_report：安全周报 - monthly_report：安全月报 - custom_report：自定义报告
type ReportCategoryRes struct {
}

func (o ReportCategoryRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportCategoryRes struct{}"
	}

	return strings.Join([]string{"ReportCategoryRes", string(data)}, " ")
}
