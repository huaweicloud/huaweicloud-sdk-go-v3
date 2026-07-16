package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOsQuotaRequest Request Object
type ShowOsQuotaRequest struct {
}

func (o ShowOsQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOsQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowOsQuotaRequest", string(data)}, " ")
}
