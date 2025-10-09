package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnGatewayJobsResponse Response Object
type ListVpnGatewayJobsResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表
	Jobs           *[]Job `json:"jobs,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVpnGatewayJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnGatewayJobsResponse struct{}"
	}

	return strings.Join([]string{"ListVpnGatewayJobsResponse", string(data)}, " ")
}
