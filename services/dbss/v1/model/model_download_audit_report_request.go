package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAuditReportRequest Request Object
type DownloadAuditReportRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 报表ID。可通过查询报表接口ID值获取。 **约束限制**： 不涉及 **取值范围**： 以查询报表接口值为准，字符长度16-64。 **默认取值**： 不涉及
	Id string `json:"id"`

	// **参数解释**： 语言区域。 **约束限制**： 仅支持取值范围约定的值 **取值范围**： - en-us : 英语 - zh-cn : 中文 **默认取值**： en-us
	Local *string `json:"local,omitempty"`
}

func (o DownloadAuditReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditReportRequest struct{}"
	}

	return strings.Join([]string{"DownloadAuditReportRequest", string(data)}, " ")
}
