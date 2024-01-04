package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobAuditResultResponse Response Object
type ShowJobAuditResultResponse struct {
	SystemAuditResult *AuditResultSystemAuditResult `json:"system_audit_result,omitempty"`

	AdminAuditResult *AuditResultAdminAuditResult `json:"admin_audit_result,omitempty"`
	HttpStatusCode   int                          `json:"-"`
}

func (o ShowJobAuditResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobAuditResultResponse struct{}"
	}

	return strings.Join([]string{"ShowJobAuditResultResponse", string(data)}, " ")
}
