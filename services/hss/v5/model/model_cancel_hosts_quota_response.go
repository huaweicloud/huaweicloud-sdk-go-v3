package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelHostsQuotaResponse Response Object
type CancelHostsQuotaResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelHostsQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelHostsQuotaResponse struct{}"
	}

	return strings.Join([]string{"CancelHostsQuotaResponse", string(data)}, " ")
}
