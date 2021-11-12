package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowQuotaDefaultsRequest struct {
}

func (o ShowQuotaDefaultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaDefaultsRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaDefaultsRequest", string(data)}, " ")
}
