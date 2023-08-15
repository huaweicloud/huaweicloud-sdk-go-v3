package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseProjectQuotaRequest Request Object
type ShowEnterpriseProjectQuotaRequest struct {
}

func (o ShowEnterpriseProjectQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectQuotaRequest", string(data)}, " ")
}
