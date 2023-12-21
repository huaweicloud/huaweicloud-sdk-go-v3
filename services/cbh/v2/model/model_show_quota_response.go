package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaResponse Response Object
type ShowQuotaResponse struct {

	// 云堡垒机实例最大创建数量。
	Quota *int32 `json:"quota,omitempty"`

	// 当前云堡垒机实例创建个数。
	QuotaUsed      *int32 `json:"quota_used,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
