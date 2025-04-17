package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaResponse Response Object
type ListQuotaResponse struct {
	DdosQuota      *QuotaResponseDdosQuota `json:"ddos_quota,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaResponse", string(data)}, " ")
}
