package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportIdReq **参数解释**: 报告ID **约束限制**: 必填 **取值范围**: 最小值0，最大值2147483647 **默认取值**: 不涉及
type ReportIdReq struct {
}

func (o ReportIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportIdReq struct{}"
	}

	return strings.Join([]string{"ReportIdReq", string(data)}, " ")
}
