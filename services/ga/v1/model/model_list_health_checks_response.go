package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthChecksResponse Response Object
type ListHealthChecksResponse struct {

	// 健康检查列表。
	HealthChecks *[]HealthCheckDetail `json:"health_checks,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHealthChecksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthChecksResponse struct{}"
	}

	return strings.Join([]string{"ListHealthChecksResponse", string(data)}, " ")
}
