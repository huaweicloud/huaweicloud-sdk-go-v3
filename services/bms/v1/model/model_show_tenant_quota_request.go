package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTenantQuotaRequest struct {
}

func (o ShowTenantQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantQuotaRequest", string(data)}, " ")
}
