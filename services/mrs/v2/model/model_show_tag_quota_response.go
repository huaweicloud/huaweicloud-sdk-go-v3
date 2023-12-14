package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagQuotaResponse Response Object
type ShowTagQuotaResponse struct {

	// 总配额大小
	TotalQuota *int32 `json:"total_quota,omitempty"`

	// 可使用配额大小
	AvailableQuota *int32 `json:"available_quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTagQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowTagQuotaResponse", string(data)}, " ")
}
