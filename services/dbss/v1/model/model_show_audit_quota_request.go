package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditQuotaRequest Request Object
type ShowAuditQuotaRequest struct {
}

func (o ShowAuditQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditQuotaRequest", string(data)}, " ")
}
