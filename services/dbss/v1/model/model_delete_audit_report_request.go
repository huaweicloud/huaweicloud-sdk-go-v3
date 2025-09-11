package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditReportRequest Request Object
type DeleteAuditReportRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 报表ID。可通过查询报表接口ID值获取。 **约束限制**： 不涉及 **取值范围**： 以查询报表接口值为准，字符长度16-64。 **默认取值**： 不涉及
	Id string `json:"id"`
}

func (o DeleteAuditReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditReportRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuditReportRequest", string(data)}, " ")
}
