package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditBackupRiskTemplatesResponse Response Object
type ListAuditBackupRiskTemplatesResponse struct {

	// 配置列表
	Templates      *[]RiskBackupTemplateBean `json:"templates,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAuditBackupRiskTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditBackupRiskTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAuditBackupRiskTemplatesResponse", string(data)}, " ")
}
