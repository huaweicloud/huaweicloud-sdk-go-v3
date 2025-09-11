package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditReportTemplatesResponse Response Object
type ListAuditReportTemplatesResponse struct {

	// os类型
	OsType *string `json:"os_type,omitempty"`

	// 模板列表
	Templates      *[]TemplateInfo `json:"templates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAuditReportTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditReportTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAuditReportTemplatesResponse", string(data)}, " ")
}
