package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQuotaResponse struct {

	// 剩余可创建云堡垒机实例个数。
	Quota *int32 `json:"quota,omitempty"`

	// 弹性公网IP个数，返回默认值1。
	EipQuota *int32 `json:"eip_quota,omitempty"`

	// IPV6ECS资源状态信息，返回默认值null。
	StatusV6 *string `json:"status_v6,omitempty"`

	// ECS资源状态信息，返回默认值null。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQuotaResponse struct{}"
	}

	return strings.Join([]string{"SearchQuotaResponse", string(data)}, " ")
}
