package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSecurityReportRequestInfo struct {

	// **参数解释**： 报告ID **约束限制**： 必填 **取值范围**： 最小值0，最大值2147483647 **默认取值**： 不涉及
	ReportId int32 `json:"report_id"`
}

func (o DeleteSecurityReportRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityReportRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteSecurityReportRequestInfo", string(data)}, " ")
}
