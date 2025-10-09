package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVpnGatewayJobsResponse Response Object
type ListP2cVpnGatewayJobsResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表
	Jobs           *[]Job `json:"jobs,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListP2cVpnGatewayJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVpnGatewayJobsResponse struct{}"
	}

	return strings.Join([]string{"ListP2cVpnGatewayJobsResponse", string(data)}, " ")
}
