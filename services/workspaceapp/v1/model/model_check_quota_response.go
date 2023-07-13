package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckQuotaResponse Response Object
type CheckQuotaResponse struct {

	// 配额是否足够true：足够 false：不足
	IsEnough *bool `json:"is_enough,omitempty"`

	// 配额剩余数量信息
	QuotaRemainder *[]QuotaRemainderData `json:"quota_remainder,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CheckQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckQuotaResponse struct{}"
	}

	return strings.Join([]string{"CheckQuotaResponse", string(data)}, " ")
}
