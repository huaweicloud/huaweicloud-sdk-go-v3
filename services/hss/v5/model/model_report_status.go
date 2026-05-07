package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportStatus **参数解释**: 报告开启状态 **取值范围**:   - opened：开启   - closed：关闭
type ReportStatus struct {
}

func (o ReportStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportStatus struct{}"
	}

	return strings.Join([]string{"ReportStatus", string(data)}, " ")
}
