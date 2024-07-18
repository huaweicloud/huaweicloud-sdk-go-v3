package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpListResponse Response Object
type BatchDeleteIpListResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	// 参数解释：请求ID。 注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListResponse", string(data)}, " ")
}
