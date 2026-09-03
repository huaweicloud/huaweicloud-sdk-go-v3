package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOpeningInfoResponse Response Object
type ShowOpeningInfoResponse struct {

	// 是否开通
	Open *bool `json:"open,omitempty"`

	// 配额是否超过
	IsQuotaExceed *bool `json:"is_quota_exceed,omitempty"`

	// 开通配额总数
	QuotaNum *int32 `json:"quota_num,omitempty"`

	// 已使用配额数量
	UsedNum *int32 `json:"used_num,omitempty"`

	// 是否付费
	IsCharge       *bool `json:"is_charge,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowOpeningInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpeningInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowOpeningInfoResponse", string(data)}, " ")
}
