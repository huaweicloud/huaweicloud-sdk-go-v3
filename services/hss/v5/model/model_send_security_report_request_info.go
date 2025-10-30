package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSecurityReportRequestInfo 发送安全报告
type SendSecurityReportRequestInfo struct {

	// **参数解释**: 报告ID **取值范围**: 字符长度10-2147483647位
	ReportId int32 `json:"report_id"`

	// **参数解释**: 报告子ID **取值范围**: 字符长度10-2147483647位
	ReportSubId int32 `json:"report_sub_id"`
}

func (o SendSecurityReportRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSecurityReportRequestInfo struct{}"
	}

	return strings.Join([]string{"SendSecurityReportRequestInfo", string(data)}, " ")
}
