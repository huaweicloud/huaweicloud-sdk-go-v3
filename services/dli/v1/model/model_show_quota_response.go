package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaResponse Response Object
type ShowQuotaResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	Quotas         *QuotaList `json:"quotas,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
