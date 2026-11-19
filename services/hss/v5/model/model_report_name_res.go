package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportNameRes **参数解释**: 报告名称 **取值范围**: 字符长度1-32位
type ReportNameRes struct {
}

func (o ReportNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportNameRes struct{}"
	}

	return strings.Join([]string{"ReportNameRes", string(data)}, " ")
}
