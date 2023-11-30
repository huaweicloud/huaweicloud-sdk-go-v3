package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthsLimitResponse Response Object
type ListBandwidthsLimitResponse struct {

	// 带宽限制列表
	EipBandwidthLimits *[]ShowTenantDict `json:"eip_bandwidth_limits,omitempty"`

	PageInfo *PageInfoDict `json:"page_info,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBandwidthsLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthsLimitResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthsLimitResponse", string(data)}, " ")
}
