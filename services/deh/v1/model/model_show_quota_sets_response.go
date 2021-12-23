package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQuotaSetsResponse struct {
	// 专属主机的配额。

	QuotaSet       *[]RespQuotaSet `json:"quota_set,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowQuotaSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaSetsResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaSetsResponse", string(data)}, " ")
}
