package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditQuotaNewRequest Request Object
type ShowAuditQuotaNewRequest struct {
}

func (o ShowAuditQuotaNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditQuotaNewRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditQuotaNewRequest", string(data)}, " ")
}
