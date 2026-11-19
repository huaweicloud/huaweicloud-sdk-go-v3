package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportNameReq **参数解释**: 报告名称 **约束限制**: 必填 **取值范围**: 字符长度1-32位 **默认取值**: 不涉及
type ReportNameReq struct {
}

func (o ReportNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportNameReq struct{}"
	}

	return strings.Join([]string{"ReportNameReq", string(data)}, " ")
}
