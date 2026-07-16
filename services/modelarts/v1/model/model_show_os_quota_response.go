package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOsQuotaResponse Response Object
type ShowOsQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOsQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOsQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowOsQuotaResponse", string(data)}, " ")
}
