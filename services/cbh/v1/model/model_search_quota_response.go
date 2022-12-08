package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQuotaResponse struct {

	// 配额
	Quota *int32 `json:"quota,omitempty"`

	// EIP配额
	EipQuota *int32 `json:"eip_quota,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// v6状态
	StatusV6       *string `json:"status_v6,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQuotaResponse struct{}"
	}

	return strings.Join([]string{"SearchQuotaResponse", string(data)}, " ")
}
