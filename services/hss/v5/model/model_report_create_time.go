package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportCreateTime **参数解释**： 报告创建时间 **取值范围**: 不涉及
type ReportCreateTime struct {
}

func (o ReportCreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportCreateTime struct{}"
	}

	return strings.Join([]string{"ReportCreateTime", string(data)}, " ")
}
