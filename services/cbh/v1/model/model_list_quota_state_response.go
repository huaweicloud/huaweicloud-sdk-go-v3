package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQuotaStateResponse struct {

	// 配额
	Quota *string `json:"quota,omitempty"`

	// 支持IPv6 ECS资源信息
	StatusV6 *string `json:"status_v6,omitempty"`

	// ECS资源状态
	Status *string `json:"status,omitempty"`

	// EIP配额信息
	EipQuota       *string `json:"eip_quota,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListQuotaStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaStateResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaStateResponse", string(data)}, " ")
}
