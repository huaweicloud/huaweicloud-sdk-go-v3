package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipPoolsResponse Response Object
type ListGeipPoolsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 全域弹性公网IP池列表
	GeipPools *[]ListGeipPools `json:"geip_pools,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGeipPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListGeipPoolsResponse", string(data)}, " ")
}
