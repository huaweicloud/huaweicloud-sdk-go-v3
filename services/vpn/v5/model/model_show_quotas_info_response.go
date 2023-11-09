package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasInfoResponse Response Object
type ShowQuotasInfoResponse struct {
	Quotas *Quotas `json:"quotas,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuotasInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasInfoResponse", string(data)}, " ")
}
